package main

import (
    "bufio"
    "flag"
    "fmt"
    "net"
    "os"
    "strconv"
    "strings"
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

func portScan(host string, ports []int) {
    var wg sync.WaitGroup

    for _, port := range ports {
        wg.Add(1)
        go scanPort(host, port, &wg)
    }

    wg.Wait()
}

func scanFromFile(fileName string, ports []int) {
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        host := scanner.Text()
        portScan(host, ports)
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
}

func parsePorts(ports string) ([]int, error) {
    var result []int
    parts := strings.Split(ports, ",")
    for _, part := range parts {
        if strings.Contains(part, "-") {
            rangeParts := strings.Split(part, "-")
            if len(rangeParts) != 2 {
                return nil, fmt.Errorf("invalid range: %s", part)
            }
            start, err := strconv.Atoi(rangeParts[0])
            if err != nil {
                return nil, fmt.Errorf("invalid port: %s", rangeParts[0])
            }
            end, err := strconv.Atoi(rangeParts[1])
            if err != nil {
                return nil, fmt.Errorf("invalid port: %s", rangeParts[1])
            }
            for i := start; i <= end; i++ {
                result = append(result, i)
            }
        } else {
            port, err := strconv.Atoi(part)
            if err != nil {
                return nil, fmt.Errorf("invalid port: %s", part)
            }
            result = append(result, port)
        }
    }
    return result, nil
}

func main() {
    host := flag.String("h", "", "Host to scan")
    fileName := flag.String("f", "", "File with hosts to scan")
    portsFlag := flag.String("p", "", "Ports to scan (e.g., 80,443 or 1-1000)")

    flag.Parse()

    if *portsFlag == "" {
        fmt.Println("Usage: ./portscanner -h <host> -p <ports> or ./portscanner -f <filename> -p <ports>")
        return
    }

    ports, err := parsePorts(*portsFlag)
    if err != nil {
        fmt.Printf("Error parsing ports: %v\n", err)
        return
    }

    if *fileName != "" {
        scanFromFile(*fileName, ports)
    } else if *host != "" {
        portScan(*host, ports)
    } else {
        fmt.Println("Usage: ./portscanner -h <host> -p <ports> or ./portscanner -f <filename> -p <ports>")
    }
}
