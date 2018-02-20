package main

import (
	"os"
	"strings"
	"testing"
)

func TestLoadAvg(t *testing.T) {
	expectedResult := "0.00 0.01 0.00 1/198 3417"
	actualResult := getFileContent("test/loadavg")
	if !strings.Contains(actualResult, expectedResult) {
		t.Error(
			"For", "test/loadavg",
			"expected", expectedResult,
			"got", actualResult,
		)
	}
}

func TestServerPortGetter(t *testing.T) {
	expectedResult := "42"
	os.Setenv("SERVER_PORT", expectedResult)
	actualResult := getServerPort()
	if actualResult != expectedResult {
		t.Error(
			"For", "test/loadavg",
			"expected", expectedResult,
			"got", actualResult,
		)
	}
}

func TestCSVHeader(t *testing.T) {
	expectedResult := "load_1, load_5, load_10, procs_running, procs_idle, ctxt_switches, uptime, last_pid, blocked, sirq_hi, sirq_timer, sirq_NetTx, sirq_NetRx, sirq_block, sirq_block_io_poll, sirq_tasklet, sirq_scheduler, sirq_hr_timer, sirq_rcu, sirq_total"
	actualResult := prepareCSVHeader()
	if !strings.Contains(actualResult, expectedResult) {
		t.Error(
			"For", "prepareCSVHeader",
			"expected", expectedResult,
			"got", actualResult,
		)
	}
}

func TestStat(t *testing.T) {
	expectedResult := "0 2.93%, 0.81%, 96.03%, 0.16%, 0.00%, 0.07%, 2.98%, 0.86%, 96.01%, 0.16%, 0.00%, 0.00%, 2.89%, 0.76%, 96.05%, 0.17%, 0.00%, 0.13%, 8109530, 3416, 0, 4409058, 5, 478252, 2038, 1736808, 13786, 0, 476620, 432032, 0, 1269517"
	actualResult := parseStat("test/stat")
	if expectedResult != actualResult {
		t.Error(
			"For", "test/stat",
			"Expected::", expectedResult,
			"Got::", actualResult,
		)
	}
}
func TestMissingFile(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("TestUserFail should have panicked!")
			}
		}()
		getFileContent("loadavg")
	}()
}
