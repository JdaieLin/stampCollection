package main

import (
	"cici/andromeda/routers"

	"github.com/astaxie/beego"
)

func main() {
	routers.InitRouters()
	beego.Run()
}
