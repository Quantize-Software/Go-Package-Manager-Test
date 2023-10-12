package backend

import (
	"fmt"
	"os"
	"os/exec"
)

func Init() {
	fmt.Println("Hello World from go main function in github package!")
	path, _ := os.Getwd()
	fmt.Println("Working Dir:", path)

	data, err := exec.Command("node", path+"/data/server.js").Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
