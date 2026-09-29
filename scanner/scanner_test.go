package scanner

import (
	"net"
	"os/exec"
	"testing"
	"time"
)

func TestGetListeningPorts(t *testing.T) {
	// Start a dummy TCP listener
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()

	// Extract the port number
	addr := l.Addr().(*net.TCPAddr)
	port := uint32(addr.Port)

	// Wait a tiny bit for the OS to report it as LISTEN
	time.Sleep(100 * time.Millisecond)

	ports, err := GetListeningPorts()
	if err != nil {
		t.Fatalf("GetListeningPorts failed: %v", err)
	}

	found := false
	for _, p := range ports {
		if p.Port == port {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected to find port %d in listening ports, but didn't", port)
	}

	// Verify sorting
	for i := 1; i < len(ports); i++ {
		if ports[i-1].Port > ports[i].Port {
			t.Errorf("ports are not sorted: %v", ports)
			break
		}
	}
}

func TestKillProcess(t *testing.T) {
	// Start a dummy process
	cmd := exec.Command("sleep", "10")
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start dummy process: %v", err)
	}
	
	pid := int32(cmd.Process.Pid)
	
	// Ensure the process is actually running
	time.Sleep(100 * time.Millisecond)
	
	err := KillProcess(pid)
	if err != nil {
		t.Errorf("KillProcess failed: %v", err)
	}
	
	// Wait for process to exit and verify it was killed
	err = cmd.Wait()
	if err == nil {
		t.Errorf("expected error from Wait() for killed process, got nil")
	}
}
