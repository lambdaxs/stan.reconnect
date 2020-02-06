package main

import (
    "fmt"
    stan_reconnect "github.com/lambdaxs/stan.reconnect"
    "github.com/nats-io/stan.go"
)


func main() {
    sc, err := stan_reconnect.Connect("c1", "server-host-1",
        stan.NatsURL("nats://49.235.146.124:4222"))
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer sc.Close()

    if err := sc.Publish("foo", []byte("hello world"));err != nil {
        fmt.Println(err.Error())
        return
    }

    if err := sc.Publish("bar", []byte("hello world!!!"));err != nil {
        fmt.Println(err.Error())
        return
    }
}
