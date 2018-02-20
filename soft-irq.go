package main

import "fmt"

//SoftIRQ struct to represnt stat values
type SoftIRQ struct {
	Hi          float64
	Timer       float64
	NetTx       float64
	NetRx       float64
	Block       float64
	BlockIOPoll float64
	Tasklet     float64
	Scheduler   float64
	HRTimer     float64
	RCU         float64
	Total       float64
}

//Print returns string representation of @SoftIRQ
func (softIRQ SoftIRQ) Print() string {
	return fmt.Sprintf("%.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f, %.f", softIRQ.Hi, softIRQ.Timer, softIRQ.NetTx, softIRQ.NetRx, softIRQ.Block,
		softIRQ.BlockIOPoll, softIRQ.Tasklet, softIRQ.Scheduler, softIRQ.HRTimer, softIRQ.RCU, softIRQ.Total)
}

//SoftIRQHeader returns metadata for @SoftIRQ
func SoftIRQHeader() string {
	return "sirq_hi, sirq_timer, sirq_NetTx, sirq_NetRx, sirq_block, sirq_block_io_poll, sirq_tasklet, sirq_scheduler, sirq_hr_timer, sirq_rcu, sirq_total"
}

//NewSoftIRQ returns SoftIRQ instance from string array
func NewSoftIRQ(statArray []string) SoftIRQ {
	softIRQ := SoftIRQ{}
	softIRQ.Hi = parseFloat(statArray[1])
	softIRQ.Timer = parseFloat(statArray[2])
	softIRQ.NetTx = parseFloat(statArray[3])
	softIRQ.NetRx = parseFloat(statArray[4])
	softIRQ.Block = parseFloat(statArray[5])
	softIRQ.BlockIOPoll = parseFloat(statArray[6])
	softIRQ.Tasklet = parseFloat(statArray[7])
	softIRQ.Scheduler = parseFloat(statArray[8])
	softIRQ.HRTimer = parseFloat(statArray[9])
	softIRQ.RCU = parseFloat(statArray[10])
	softIRQ.Total = parseFloat(statArray[11])
	return softIRQ
}
