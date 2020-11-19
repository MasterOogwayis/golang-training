package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var str string = "{\"servers\":1[{\"serverName\":\"Shanghai_VPN\",\"serverIP\":\"127.0.0.1\"},{\"serverName\":\"Beijing_VPN\",\"serverIP\":\"127.0.0.2\"}]}"
	var s Serverslice
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Errorf("转换出错 %s", err)
	}
	fmt.Println(s)



}



type Server struct {
	ServerName string
	ServerIP   string
}
type Serverslice struct {
	Servers []Server
}