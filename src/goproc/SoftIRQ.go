package goproc

import "fmt"

//SoftIRQ struct to represnt stat values
type SoftIRQ struct {
	Hi        float64
	Timer     float64
	NetTx     float64
	NetRx     float64
	Block     float64
	Poll      float64
	Tasklet   float64
	Scheduler float64
	HRTimer   float64
	RCU       float64
	RCU2      float64
}

//Print methof for @SoftIRQ
func (softIRQ SoftIRQ) Print() string {
	return fmt.Sprintf("%.f %.f %.f %.f %.f %.f %.f %.f %.f %.f %.f", softIRQ.Hi, softIRQ.Timer, softIRQ.NetTx, softIRQ.NetRx, softIRQ.Block,
		softIRQ.Poll, softIRQ.Tasklet, softIRQ.Scheduler, softIRQ.HRTimer, softIRQ.RCU, softIRQ.RCU2)
}
