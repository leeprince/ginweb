package main

import (
	_ "ginweb-router-v3/app"
	"ginweb-router-v3/routers"
)

func main() {
	routers.IniterRouter().Run()
}
