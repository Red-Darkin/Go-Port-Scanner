package main

import (
    "fmt"
    "net"
    "os"
    "strconv"
    "sync"
    "time"
)

const (
    SYN_TIMEOUT = 5 * time.Second
)

func scanPort(host string, port int, wg *sync.WaitGroup) {
    defer wg.Done()

    target := fmt.Sprintf("%s:%d", host, port)
    conn, err := net.DialTimeout("tcp", target, SYN_TIMEOUT)
    if err != nil {
        return // Puerto cerrado o filtrado
    }
    conn.Close()
    fmt.Printf("[*] Puerto %d abierto en %s\n", port, host)
}

func portScan(host string, startPort, endPort int) {
    var wg sync.WaitGroup

    for port := startPort; port <= endPort; port++ {
        wg.Add(1)
        go scanPort(host, port, &wg)
    }

    wg.Wait()
}

func main() {
    if len(os.Args) < 4 {
        fmt.Println("Uso: ./portscanner <host> <inicio_puerto> <fin_puerto>")
        return
    }

    host := os.Args[1]
    startPort, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Error: Inicio de puerto no válido")
        return
    }

    endPort, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println("Error: Fin de puerto no válido")
        return
    }

    portScan(host, startPort, endPort)
}
