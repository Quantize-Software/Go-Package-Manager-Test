package backend

import (
	"embed"
	"fmt"
	"os/exec"
)

//go:embed data
var f embed.FS

func Init() {
	fmt.Println("Hello World from go main function in github package!")

	data, err := exec.Command("node", "server.js").Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
