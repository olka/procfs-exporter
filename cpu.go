package main

import (
	"fmt"
)

//CPU stats data structure
type CPU struct {
	Title   string
	Freq    int
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
func (cpu CPU) Print(totalCPU float64) string {
	return fmt.Sprintf("%.2f%%, %.2f%%, %.2f%%, %.2f%%, %.2f%%, %.2f%%, ", ((cpu.User+cpu.Niced)*100)/totalCPU, (cpu.System*100)/totalCPU, (cpu.Idle*100)/totalCPU, (cpu.WaitIO*100)/totalCPU, (cpu.IRQ*100)/totalCPU, (cpu.SoftIRQ*100)/totalCPU)
}

//Header returns metadata information of @CPU
func (cpu CPU) Header() string {
	fields := []string{"_user, ", "_freq, ", "_system, ", "_idle, ", "_wait_io, ", "_irq, ", "_sirq, "}
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
	return CPU{cpu.Title, (cpu.Freq - other.Freq), (cpu.User - other.User), (cpu.Niced - other.Niced),
		(cpu.System - other.System), (cpu.Idle - other.Idle), (cpu.WaitIO - other.WaitIO),
		(cpu.IRQ - other.IRQ), (cpu.SoftIRQ - other.SoftIRQ), (cpu.Total - other.Total)}
}

//NewCPU returns CPU instance from string array
func NewCPU(freq int, statArray []string, i int) CPU {
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
