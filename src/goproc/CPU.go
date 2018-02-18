package goproc

import "fmt"

//CPU stats data structure
type CPU struct {
	Title   string
	User    float64
	Niced   float64
	System  float64
	Idle    float64
	WaitIO  float64
	IRQ     float64
	SoftIRQ float64
	Total   float64
}

//Print returns string representation of @CPU struct in format title us% sys% idle% iowait% hard_irq% soft_irq%
func (cpu CPU) Print(totalCPU float64) string {
	return fmt.Sprintf("\n%v: %.2f, %.2f, %.2f, %.2f, %.2f, %.2f ", cpu.Title, ((cpu.User+cpu.Niced)*100)/totalCPU, (cpu.System*100)/totalCPU, (cpu.Idle*100)/totalCPU, (cpu.WaitIO*100)/totalCPU, (cpu.IRQ*100)/totalCPU, (cpu.SoftIRQ*100)/totalCPU)
}

//Subtract values of one @CPU from another
func (cpu CPU) Subtract(other CPU) CPU {
	return CPU{cpu.Title, (cpu.User - other.User), (cpu.Niced - other.Niced),
		(cpu.System - other.System), (cpu.Idle - other.Idle), (cpu.WaitIO - other.WaitIO),
		(cpu.IRQ - other.IRQ), (cpu.SoftIRQ - other.SoftIRQ), (cpu.Total - other.Total)}
}
