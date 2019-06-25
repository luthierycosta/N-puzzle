package main

import (
	"fmt"
	"net"
	"encoding/json"
)
func jsonHandle() {
	obj := struct {
		Begin string `json:"begin"`
		Result string `json:"result"`
	}{
		Begin: begin.ToString(),
		Result: fmt.Sprint(result),
	}
	toNode, _ := json.Marshal(obj)
	conn, err := net.Dial("udp", ":8080")
	if err != nil {
		conn.Write([]byte(toNode))
	}
}
