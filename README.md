# campstate

```bash
go get github.com/drbh/campaign-state-mgmt
```

```bash
go run examples/test.go 
```



```golang
package main

import (
    "fmt"
    campstate "github.com/drbh/campaign-state-mgmt"
)

func onComplete(t campstate.Pending) {
    fmt.Println("Send: ", t.Msg)
    if t.Fail == "" {
        fmt.Println("-Terminate")
    }

}

func main() {

    var config = `[
    {
        "id": "message1",
        "msg": "Will you respond in time to get first class tickets!?!",

        "pass": "message2",
        "fail": "message3"
    },
    {
        "id": "message2",
        "msg": "Yes! you answered in time! here is a link to you tickets -- Do you want a reminder?",

        "pass": "message4",
        "fail": "message5"
    },
    {
        "id": "message3",
        "msg": "Womp womp womp, maybe next time!"
    },
    {
        "id": "message4",
        "msg": "Hey don't forget to come to the show! It's in an hour"
    },
    {
        "id": "message5",
        "msg": "Removing you from the contact list"
    }]
    `

    camp := campstate.Campaign{
        Title:    "Ticket Promotion",
        Json:     config,
        Callback: onComplete,
    }
    camp.Build()

    err := camp.FSM.Event("message1-pass")
    if err != nil {
        fmt.Println(err)
    }

    err = camp.FSM.Event("message2-pass")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("\n\nDifferent Sitch")
    camp.Build()

    err = camp.FSM.Event("message1-pass")
    if err != nil {
        fmt.Println(err)
    }

    err = camp.FSM.Event("message2-fail")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("\n\nDifferent Sitch")
    camp.Build()

    err = camp.FSM.Event("message1-fail")
    if err != nil {
        fmt.Println(err)
    }

}
```

### Response 

```bash
Send:  Will you respond in time to get first class tickets!?!
Send:  Yes! you answered in time! here is a link to you tickets -- Do you want a reminder?
Send:  Hey don't forget to come to the show! It's in an hour
-Terminate


Different Sitch
Send:  Will you respond in time to get first class tickets!?!
Send:  Yes! you answered in time! here is a link to you tickets -- Do you want a reminder?
Send:  Removing you from the contact list
-Terminate


Different Sitch
Send:  Will you respond in time to get first class tickets!?!
Send:  Womp womp womp, maybe next time!
-Terminate
```
