package main

import (
	"fmt"
	"github.com/stlf2004/test/pkgname/pkg01"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println(pkg01.Name)
	beego.Run()
}
