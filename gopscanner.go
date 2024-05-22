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
        return 
    }
    conn.Close()
    fmt.Printf("[*] Open Port %s:%d\n", host, port)
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
        fmt.Println("Usage: ./portscanner <host> <start_port> <end_port>")
        return
    }

    host := os.Args[1]
    startPort, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Error: Invalid startPort")
        return
    }

    endPort, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println("Error: Invalid endPort")
        return
    }

    portScan(host, startPort, endPort)
}
