package gpu

import (
	"golang.org/x/sys/cpu"
)

func GetCPUCapability() CPUCapability {
	if cpu.X86.HasAVX2 {
		return CPUCapabilityAVX2
	}
	if cpu.X86.HasAVX {
		return CPUCapabilityAVX
	}
	// else LCD
	return CPUCapabilityNone
}

func GetARM64CpuCapability() CPUCapability {
	switch cpu.extractBits(cpu.getisar1(), 55, 52) {
	case 1:
		return CPUCapabilityI8MM
	}
}
