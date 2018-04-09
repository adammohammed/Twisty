package main

import (
        "bufio"
        "log"
        "net"
        "github.com/jacobsa/go-serial/serial"
)

func main() {

        // Dial Server
        conn, err := net.Dial("tcp", "adammohammed.org:9090")
        if err != nil {
                log.Fatalf("Couldn't connect to tcp server\n")
        }
        defer conn.Close()

        // Set up options.
        options := serial.OpenOptions{
                PortName: "/dev/ttyACM0",
                BaudRate: 9600,
                DataBits: 8,
                StopBits: 1,
                MinimumReadSize: 4,
        }

        log.Println("Successfully Connected")

        // Open the port.
        port, err := serial.Open(options)
        if err != nil {
                log.Fatalf("serial.Open: %v", err)
         }

        // Make sure to close it later.
        defer port.Close()


        // Listen for messages
        for {
                b := make([]byte,8)
                bufio.NewReader(conn).Read(b)
                log.Printf("Received character %c\n", b[0])
                _, err := port.Write(b)
                if err != nil {
                        log.Fatalf("port.Write: %v", err)
                }
        }

}
