package main
import (
    "fmt"
    "net"
)

func main() {
    conn, err := net.Dial("udp", "127.0.0.1:8080")
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
    fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
    conn.Close()
}
