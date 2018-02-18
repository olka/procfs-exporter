package main

import "testing"

func TestLoadAvg(t *testing.T) {
	expectedResult := "0.00 0.01 0.00 1/198 3417\n"
	actualResult := getFileContent("test/loadavg")
	if expectedResult == actualResult {
		t.Error(
			"For", "test/loadavg",
			"expected", expectedResult,
			"got", actualResult,
		)
	}
}

func TestStat(t *testing.T) {
	expectedResult := `cpu: 2.93, 0.81, 96.03, 0.16, 0.00, 0.07
						cpu0: 2.98, 0.86, 96.01, 0.16, 0.00, 0.00
						cpu1: 2.89, 0.76, 96.05, 0.17, 0.00, 0.13
						ctxt: 8109530
						uptime: 36h54m31s
						processes: 3416
						procs_running: 1
						procs_blocked: 0
						softirq: 4409058 5 478252 2038 1736808 13786 0 476620 432032 0 1269517`
	actualResult := parseStat(getFileContent("test/stat"))
	if expectedResult == actualResult {
		t.Error(
			"For", "test/stat",
			"expected", expectedResult,
			"got", actualResult,
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
