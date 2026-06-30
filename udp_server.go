package main

import (
    "fmt"
    "net"
)

func main() {
    // 1. Слушаем UDP-адрес
    addr, err := net.ResolveUDPAddr("udp", ":8080")
    if err != nil {
        panic(err)
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    fmt.Println("UDP сервер слушает на порту 8080")

    buf := make([]byte, 1024)

    for {
        // 2. Читаем пакет
        n, clientAddr, err := conn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("Ошибка чтения:", err)
            continue
        }

        message := string(buf[:n])
        fmt.Printf("Получено от %s: %s\n", clientAddr, message)

        // 3. Отвечаем
        response := []byte("Привет от Go сервера!")
        _, err = conn.WriteToUDP(response, clientAddr)
        if err != nil {
            fmt.Println("Ошибка отправки:", err)
        }
    }
}
