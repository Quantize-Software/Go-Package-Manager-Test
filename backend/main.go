package backend

import (
	"embed"
	"fmt"
	"os/exec"
)

//go:embed data
var data embed.FS

func Init() {
	fmt.Println("Hello World from go main function in github package!")
	contents, err := data.ReadFile("data/server.js")
	if err != nil {
		panic(err)
	}

	data, err := exec.Command("node", "-e", string(contents)).Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
