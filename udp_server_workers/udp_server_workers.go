package main

import (
	"fmt"
	"net"
	"runtime"
)

const (
	udpServerAddress = ":3200"
	maxPacketSize    = 1024
)

type udpJob struct {
	message    string
	clientAddr *net.UDPAddr
}

func workerCount() int {
	return runtime.NumCPU()
}

func workerResponse(workerID int) []byte {
	return []byte(fmt.Sprintf("я воркер № %d", workerID))
}

func worker(id int, jobs <-chan udpJob, conn *net.UDPConn) {
	for job := range jobs {
		fmt.Printf("Воркер № %d получил от %s: %s\n", id, job.clientAddr, job.message)

		if _, err := conn.WriteToUDP(workerResponse(id), job.clientAddr); err != nil {
			fmt.Println("Ошибка отправки:", err)
		}
	}
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", udpServerAddress)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	workers := workerCount()
	jobs := make(chan udpJob, workers)

	for id := 1; id <= workers; id++ {
		go worker(id, jobs, conn)
	}

	fmt.Printf("UDP сервер слушает на порту 3200, воркеров: %d\n", workers)

	for {
		buf := make([]byte, maxPacketSize)
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			continue
		}

		message := string(buf[:n])
		fmt.Printf("Получено от %s: %s\n", clientAddr, message)

		jobs <- udpJob{
			message:    message,
			clientAddr: clientAddr,
		}
	}
}
