package main

import (
	"net"
	"testing"
	"time"
)

func TestClientMessageUsesMessageNumber(t *testing.T) {
	if got, want := clientMessage(12), "сообщение №12"; got != want {
		t.Fatalf("clientMessage(12) = %q, want %q", got, want)
	}
}

func TestSendUDPMessageReturnsServerResponse(t *testing.T) {
	serverConn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 0,
	})
	if err != nil {
		t.Fatalf("listen udp server: %v", err)
	}
	defer serverConn.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)

		buf := make([]byte, 128)
		n, clientAddr, err := serverConn.ReadFromUDP(buf)
		if err != nil {
			t.Errorf("read udp request: %v", err)
			return
		}

		if got, want := string(buf[:n]), "сообщение №2"; got != want {
			t.Errorf("request message = %q, want %q", got, want)
			return
		}

		if _, err := serverConn.WriteToUDP([]byte("я воркер № 4"), clientAddr); err != nil {
			t.Errorf("write udp response: %v", err)
		}
	}()

	response, err := sendUDPMessage(serverConn.LocalAddr().String(), "сообщение №2", time.Second)
	if err != nil {
		t.Fatalf("send udp message: %v", err)
	}

	if got, want := response, "я воркер № 4"; got != want {
		t.Fatalf("response = %q, want %q", got, want)
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("udp server did not receive the request")
	}
}
