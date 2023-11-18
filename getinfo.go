package main

import (
	"encoding/json"
	"fmt"

	"github.com/elementsproject/glightning/glightning"
)

const (
	dirRPC  = "<YOUR_PATH>"
	fileRPC = "lightning-rpc"
)

func main() {
	lightning := glightning.NewLightning()
	err := lightning.StartUp(fileRPC, dirRPC)
	if err != nil {
		fmt.Printf("StartUp: %v", err)
		return
	}

	res, err := lightning.GetInfo()
	if err != nil {
		fmt.Printf("GetInfo: %v\n", err)
		return
	}
	jsonBytes, err := json.Marshal(res)
	if err == nil {
		fmt.Printf("%s\n\n", string(jsonBytes))
	}
}
