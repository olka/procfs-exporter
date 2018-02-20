package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/nytimes/gziphandler"
)

// DefaultPort is the default port to use if once is not specified by the SERVER_PORT environment variable
const DefaultPort = "9100"

//LoadAvgFileName file name for loadavg metrics
const LoadAvgFileName = "/proc/loadavg"

//StatFileName file name for stat metrics
var StatFileName = "/proc/stat"

//CPUFreqFileName represents path to file with current CPU frequency
var CPUFreqFileName string

var cpuState [10]CPU
var header string

func initialize() {
	getFrequencyFileName := func(cpuTitle string) string {
		prefix := "/sys/devices/system/cpu/" + cpuTitle + "/cpufreq/"
		if _, err := os.Stat(prefix + "scaling_cur_freq"); err == nil {
			return prefix + "scaling_cur_freq"
		}
		if _, err := os.Stat(prefix + "cpuinfo_cur_freq"); err == nil {
			return prefix + "cpuinfo_cur_freq"
		}
		return ""
	}

	parseStat(StatFileName)
	header = prepareCSVHeader()
	CPUFreqFileName = getFrequencyFileName(cpuState[1].Title)
}

func getServerPort() string {
	port := os.Getenv("SERVER_PORT")
	if port != "" {
		return port
	}
	return DefaultPort
}

func getFileContent(fileName string) string {
	res, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(res))
}

func parseLoadAvg(loadAvgFileName string) string {
	loadAvg := getFileContent(loadAvgFileName)
	splitter := func(r rune) bool {
		return unicode.IsSpace(r) || r == '/'
	}
	loadAvgArray := strings.FieldsFunc(loadAvg, splitter)
	loadAvgArray = loadAvgArray[:len(loadAvgArray)-1]
	return strings.Join(loadAvgArray, ", ") + ","
}

func parseBootTime(btime string) string {
	i, err := strconv.ParseInt(btime, 10, 64)
	if err != nil {
		panic(err)
	}
	timestamp := time.Since(time.Unix(i, 0)).Round(time.Second)
	return timestamp.String()
}

//https://www.kernel.org/doc/Documentation/cpu-freq/user-guide.txt
func getFrequency(cpuTitle string) int64 {
	if CPUFreqFileName != "" {
		return parseInt(getFileContent(CPUFreqFileName))
	}
	return 0
}

func parseStat(statFileName string) string {
	stat := getFileContent(statFileName)
	var result []string
	statArray := strings.FieldsFunc(stat, unicode.IsSpace)

	for i, statEntry := range statArray {
		switch {
		case strings.HasPrefix(statEntry, "cpu"):
			freq := getFrequency(statEntry)
			j := i / 11 //count of cpu related elements on one line: http://man7.org/linux/man-pages/man5/proc.5.html
			cpuCurrentVal := NewCPU(freq, statArray, i)
			result = append(result, cpuCurrentVal.Subtract(cpuState[j]).Print())
			cpuState[j] = cpuCurrentVal
		case statEntry == "ctxt":
			result = append(result, statArray[i+1])
		case statEntry == "btime":
			result = append(result, ", "+parseBootTime(statArray[i+1]))
		case statEntry == "processes":
			result = append(result, ", "+statArray[i+1])
		case statEntry == "procs_blocked":
			result = append(result, ", "+statArray[i+1])
		case statEntry == "softirq":
			result = append(result, ", "+NewSoftIRQ(statArray[i:]).Print())
		}
	}
	return strings.Join(result, "")
}

func prepareCSVHeader() string {
	var result bytes.Buffer
	result.WriteString("load_1, load_5, load_10, procs_running, procs_idle, ")
	for _, val := range cpuState {
		result.WriteString(val.Header())
	}
	result.WriteString("ctxt_switches, uptime, last_pid, blocked, ")
	result.WriteString(SoftIRQHeader())
	result.WriteString("\n")
	return result.String()
}

// ProcHandler sends metrics from proc fs
func ProcHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	var result bytes.Buffer
	result.WriteString(header)
	result.WriteString(parseLoadAvg(LoadAvgFileName))
	result.WriteString(parseStat(StatFileName))
	io.WriteString(writer, result.String())
}

func main() {
	log.Println("Starting proc agent on port " + getServerPort())
	initialize()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Use /metrics url for host metrics`))
	})
	http.Handle("/metrics", gziphandler.GzipHandler(http.HandlerFunc(ProcHandler)))

	err := http.ListenAndServe(":"+getServerPort(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
