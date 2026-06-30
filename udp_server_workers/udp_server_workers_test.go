package main

import (
	"net"
	"runtime"
	"testing"
	"time"
)

func TestWorkerCountUsesAvailableCPUCount(t *testing.T) {
	if got, want := workerCount(), runtime.NumCPU(); got != want {
		t.Fatalf("workerCount() = %d, want %d", got, want)
	}
}

func TestWorkerRespondsWithWorkerNumber(t *testing.T) {
	serverConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 0,
	})
	if err != nil {
		t.Fatalf("listen server udp: %v", err)
	}
	defer serverConn.Close()

	clientConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 0,
	})
	if err != nil {
		t.Fatalf("listen client udp: %v", err)
	}
	defer clientConn.Close()

	jobs := make(chan udpJob, 1)
	go worker(3, jobs, serverConn)

	jobs <- udpJob{
		message:    "ping",
		clientAddr: clientConn.LocalAddr().(*net.UDPAddr),
	}

	buf := make([]byte, 128)
	if err := clientConn.SetReadDeadline(time.Now().Add(time.Second)); err != nil {
		t.Fatalf("set read deadline: %v", err)
	}

	n, _, err := clientConn.ReadFromUDP(buf)
	if err != nil {
		t.Fatalf("read worker response: %v", err)
	}

	if got, want := string(buf[:n]), "я воркер № 3"; got != want {
		t.Fatalf("response = %q, want %q", got, want)
	}
}
