package schedule

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func SwagInit() {
	//TODO: 修好这个bug
	command := exec.Command("swag", "init")

	var out bytes.Buffer

	command.Stdout = &out

	command.Stderr = os.Stderr

	err := command.Start()
	if err != nil {
		log.Println("swag init error ", err.Error())
	} else {
		fmt.Println(out.String())
	}
	fmt.Println("swag init finish..")
}
