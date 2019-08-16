package campstate

import (
    "encoding/json"
    "fmt"
    "github.com/looplab/fsm"
)

type Campaign struct {
    Title    string
    Json     string
    Callback func(Pending)
    FSM      *fsm.FSM
}

type Pending struct {
    ID   string `json:"id"`
    Msg  string `json:"msg"`
    Pass string `json:"pass,omitempty"`
    Fail string `json:"fail,omitempty"`
}

func (camp *Campaign) Build() {

    var pfadasgergregwergewq []Pending
    var fff []fsm.EventDesc

    callbizacks := make(map[string]fsm.Callback)

    if err := json.Unmarshal([]byte(camp.Json), &pfadasgergregwergewq); err != nil {
        panic(err)
    }

    for _, v := range pfadasgergregwergewq {
        var dst string
        if v.Pass != "" {
            dst = v.Pass
            item := fsm.EventDesc{
                v.ID + "-pass",
                []string{v.ID},
                dst,
            }
            fff = append(fff, item)
        }
        if v.Fail != "" {
            dst = v.Fail
            item := fsm.EventDesc{
                v.ID + "-fail",
                []string{v.ID},
                dst,
            }
            fff = append(fff, item)

        }

        callbizacks[v.ID] = func(e *fsm.Event) {
            var t Pending
            for _, b := range pfadasgergregwergewq {
                if b.ID == e.Dst {
                    t = b
                    break
                }
            }
            camp.Callback(t)
        }

    }

    camp.FSM = fsm.NewFSM(
        "message1",
        fff,
        callbizacks,
    )

    var t Pending
    for _, b := range pfadasgergregwergewq {
        if b.ID == camp.FSM.Current() {
            t = b
            break
        }
    }
    camp.Callback(t)
}

func onComplete(t Pending) {
    fmt.Println("Send: ", t.Msg)
    if t.Fail == "" {
        fmt.Println("-Terminate")
    }

}
