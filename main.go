package main

import (
	"hasdsd.cn/op-dl-server/router"
	"hasdsd.cn/op-dl-server/schedule"
)

func main() {
	schedule.SwagInit()
	r := router.Router()
	r.Run()
}
