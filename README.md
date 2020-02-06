# stan.reconnect

- reconnect plugin for stan.go

### example

```go
package main

import (
    "fmt"
    "github.com/nats-io/stan.go"
    stan_reconnect "github.com/lambdaxs/stan.reconnect"
)

func main() {
    sc, err := stan_reconnect.Connect("c1", "server-host-2", stan.NatsURL("nats://49.235.146.124:4222"))
        if err != nil {
            fmt.Println("err:" + err.Error())
            return
        }
    
    
        startOpt := stan.DeliverAllAvailable()
        _, err = sc.QueueSubscribe("foo", "consumer_group_1", func(msg *stan.Msg) {
            fmt.Printf("receive msg：%s time:%d seq:%d\n", string(msg.Data), msg.Timestamp, msg.Sequence)
            msg.Ack()
        }, startOpt, stan.DurableName("consumer_1"), stan.SetManualAckMode())
        if err != nil {
            fmt.Println(err.Error())
            return
        }
}

```




