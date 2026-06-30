package main

import (
	"fmt"
	"net"
	"time"
)

const (
	clientServerAddress = "127.0.0.1:3200"
	clientSendInterval  = 3 * time.Second
	clientTimeout       = 2 * time.Second
	clientMaxPacketSize = 1024
)

func clientMessage(number int) string {
	return fmt.Sprintf("сообщение №%d", number)
}

func sendUDPMessage(serverAddress string, message string, timeout time.Duration) (string, error) {
	serverAddr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		return "", err
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return "", err
	}

	if _, err := conn.Write([]byte(message)); err != nil {
		return "", err
	}

	buf := make([]byte, clientMaxPacketSize)
	n, err := conn.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

func main() {
	fmt.Printf("UDP клиент отправляет сообщения на %s раз в 3 секунды\n", clientServerAddress)

	for messageNumber := 1; ; messageNumber++ {
		message := clientMessage(messageNumber)
		fmt.Println("Отправлено:", message)

		response, err := sendUDPMessage(clientServerAddress, message, clientTimeout)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Ответ:", response)
		}

		time.Sleep(clientSendInterval)
	}
}
