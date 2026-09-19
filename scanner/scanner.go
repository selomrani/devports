package scanner

import (
	"sort"

	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

type PortInfo struct {
	Port    uint32
	PID     int32
	Process string
	Memory  float32
}

func GetListeningPorts() ([]PortInfo, error) {
	conns, err := net.Connections("tcp")
	if err != nil {
		return nil, err
	}

	portMap := make(map[uint32]PortInfo)

	for _, c := range conns {
		if c.Status == "LISTEN" && c.Laddr.Port != 0 {
			pi := PortInfo{
				Port: c.Laddr.Port,
				PID:  c.Pid,
			}
			if c.Pid > 0 {
				proc, err := process.NewProcess(c.Pid)
				if err == nil {
					name, _ := proc.Name()
					pi.Process = name
					mem, _ := proc.MemoryPercent()
					pi.Memory = mem
				}
			} else {
				pi.Process = "unknown/root"
			}
			
			// Ignore IPv6 duplicate of IPv4
			if existing, exists := portMap[pi.Port]; !exists || (existing.Process == "unknown/root" && pi.Process != "unknown/root") {
				portMap[pi.Port] = pi
			}
		}
	}

	var results []PortInfo
	for _, p := range portMap {
		results = append(results, p)
	}

	// Sort by port number
	sort.Slice(results, func(i, j int) bool {
		return results[i].Port < results[j].Port
	})

	return results, nil
}

func KillProcess(pid int32) error {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return err
	}
	return proc.Kill()
}
