package main

import (
	"fmt"
)

//CPU stats data structure
type CPU struct {
	Title   string
	Freq    int64
	User    float64
	Niced   float64
	System  float64
	Idle    float64
	WaitIO  float64
	IRQ     float64
	SoftIRQ float64
	Total   float64
}

//Print returns string representation of @CPU
func (cpu CPU) Print() string {
	return fmt.Sprintf("%v, %.2f%%, %.2f%%, %.2f%%, %.2f%%, %.2f%%, %.2f%%, ", cpu.Freq, ((cpu.User+cpu.Niced)*100)/cpu.Total, (cpu.System*100)/cpu.Total, (cpu.Idle*100)/cpu.Total, (cpu.WaitIO*100)/cpu.Total, (cpu.IRQ*100)/cpu.Total, (cpu.SoftIRQ*100)/cpu.Total)
}

//Header returns metadata information of @CPU
func (cpu CPU) Header() string {
	fields := []string{"_freq, ", "_user, ", "_system, ", "_idle, ", "_wait_io, ", "_irq, ", "_sirq, "}
	var res string
	for _, val := range fields {
		if cpu.Title != "" {
			res += cpu.Title + val
		} else {
			break
		}
	}
	return res
}

//Subtract values of one @CPU from another
func (cpu CPU) Subtract(other CPU) CPU {
	return CPU{cpu.Title, cpu.Freq, (cpu.User - other.User), (cpu.Niced - other.Niced),
		(cpu.System - other.System), (cpu.Idle - other.Idle), (cpu.WaitIO - other.WaitIO),
		(cpu.IRQ - other.IRQ), (cpu.SoftIRQ - other.SoftIRQ), (cpu.Total - other.Total)}
}

//NewCPU returns CPU instance from string array
func NewCPU(freq int64, statArray []string, i int) CPU {
	cpu := CPU{}
	cpu.Title = statArray[i]
	cpu.Freq = freq
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
