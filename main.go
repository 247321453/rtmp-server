package main

import (
	"./rtmp"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var port = 1935
	var err error
	if len(os.Args) > 1 {
		port, err = strconv.Atoi(os.Args[1])
	}

	if err != nil {
		fmt.Println("args error: %s", err)
		return
	}
	rtmp.LogLevel(1)
	rtmp.SimpleServer(port)
}
