package init

import (
	"fmt"
	"os/exec"
)

func init() {
	fmt.Println("Hello World from go main function in github package!")
	data, err := exec.Command("node", "data/server.js").Output()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
