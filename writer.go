package main

import (
    "context"
    "github.com/segmentio/kafka-go"
    "time"
)

func main() {
    conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "nima", 0)
    _ = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
    batch := conn.ReadBatch(10e3, 10e6)

    res := make([]byte, 10e3)

    for {
        _, err := batch.Read(res)
        if err != nil {
            break
        }
        println(string(res))
    }

    _ = batch.Close()
    _ = conn.Close()
}
