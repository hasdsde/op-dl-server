package schedule

import (
	"fmt"
	"log"
	"os/exec"
)

func SwagInit() {
	command := exec.Command("swag", "init")

	res, err := command.CombinedOutput()
	if err != nil {
		log.Println("swag init error ", err.Error())
	} else {
		fmt.Println(res)
	}
	fmt.Println("swag init finish..")
}
