package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type wsMsg struct {
	Type string `json:"type"`
	Cmd  string `json:"cmd"`
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
}

func main() {
	// strings     反射 检查类型
	s := "hello world! go language"
	s_split := strings.Split(s, " ")
	s_join := strings.Join(s_split, "")
	fmt.Println(reflect.TypeOf(s_split))
	fmt.Println(s_join)

	msgObj := wsMsg{}

	var wsData []byte

	if err := json.Unmarshal(wsData, &msgObj); err != nil {
		fmt.Print("错误1  ReceiveWsMsg")
		logrus.WithError(err).WithField("wsData", string(wsData)).Error("unmarshal websocket message failed")
	}
}
