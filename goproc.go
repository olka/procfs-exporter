package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"unicode"

	"./src/goproc"
	"github.com/nytimes/gziphandler"
)

// DefaultPort is the default port to use if once is not specified by the SERVER_PORT environment variable
const DefaultPort = "9100"

var cpuState [9]goproc.CPU
var deltaCPUState [9]goproc.CPU

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
	return string(res)
}

func parseFloat(s string) float64 {
	res, err := strconv.ParseFloat(s, 64)
	if err != nil {
		debug.PrintStack()
		log.Fatal(err)
	}
	return math.Floor(res + .5)
}

func parseBootTime(btime string) string {
	i, err := strconv.ParseInt(btime, 10, 64)
	if err != nil {
		panic(err)
	}
	timestamp := time.Since(time.Unix(i, 0)).Round(time.Second)
	return timestamp.String()
}

func parseCPU(statArray []string, i int) goproc.CPU {
	cpu := goproc.CPU{}
	cpu.Title = statArray[i]
	cpu.User = parseFloat(statArray[i+1])
	cpu.Niced = parseFloat(statArray[i+2])
	cpu.System = parseFloat(statArray[i+3])
	cpu.Idle = parseFloat(statArray[i+4])
	cpu.WaitIO = parseFloat(statArray[i+5])
	cpu.IRQ = parseFloat(statArray[i+6])
	cpu.SoftIRQ = parseFloat(statArray[i+7])
	cpu.Total = cpu.User + cpu.Niced + cpu.System + cpu.Idle + cpu.WaitIO + cpu.IRQ + cpu.SoftIRQ
	return cpu
}

func parseSoftIRQ(statArray []string) goproc.SoftIRQ {
	softIRQ := goproc.SoftIRQ{}
	softIRQ.Hi = parseFloat(statArray[1])
	softIRQ.Timer = parseFloat(statArray[2])
	softIRQ.NetTx = parseFloat(statArray[3])
	softIRQ.NetRx = parseFloat(statArray[4])
	softIRQ.Block = parseFloat(statArray[5])
	softIRQ.Poll = parseFloat(statArray[6])
	softIRQ.Tasklet = parseFloat(statArray[7])
	softIRQ.Scheduler = parseFloat(statArray[8])
	softIRQ.HRTimer = parseFloat(statArray[9])
	softIRQ.RCU = parseFloat(statArray[10])
	softIRQ.RCU2 = parseFloat(statArray[11])
	return softIRQ
}

func parseStat(stat string) string {
	var result []string
	statArray := strings.FieldsFunc(stat, unicode.IsSpace)

	for i, val := range statArray {
		switch {
		case strings.HasPrefix(val, "cpu"):
			j := i / 11 //count of cpu related elements on one line: http://man7.org/linux/man-pages/man5/proc.5.html
			deltaCPUState[j] = parseCPU(statArray, i).Subtract(cpuState[j])
			result = append(result, deltaCPUState[j].Print(deltaCPUState[j].Total))
			cpuState[j] = parseCPU(statArray, i)
		case val == "ctxt":
			result = append(result, "\nctxt: "+statArray[i+1])
		case val == "btime":
			result = append(result, "\nuptime: "+parseBootTime(statArray[i+1]))
		case val == "processes":
			result = append(result, "\nprocesses: "+statArray[i+1])
		case val == "procs_running":
			result = append(result, "\nprocs_running: "+statArray[i+1])
		case val == "procs_blocked":
			result = append(result, "\nprocs_blocked: "+statArray[i+1])
		case val == "softirq":
			result = append(result, "\nsoftirq: "+parseSoftIRQ(statArray[i:]).Print())
		}
	}
	return strings.Join(result, "")
}

// ProcHandler sends metrics from proc fs
func ProcHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain")
	var result bytes.Buffer
	result.WriteString("loadavg: " + getFileContent("/proc/loadavg"))
	result.WriteString(parseStat(getFileContent("/proc/stat")))

	io.WriteString(writer, result.String())
}

func main() {
	log.Println("Starting proc agent on port " + getServerPort())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Use /metrics url to see host metrics`))
	})
	http.Handle("/metrics", gziphandler.GzipHandler(http.HandlerFunc(ProcHandler)))

	err := http.ListenAndServe(":"+getServerPort(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
