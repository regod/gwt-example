package main

import (
	"github.com/regod/gwt"
	"github.com/regod/gwt-example/user"
)

func main() {
	app := gwt.New()

	app.AddRoute("POST", "/user/create/", user.Create, nil)
	app.AddRoute("POST", "/user/update_phone/:id", user.UpdatePhone, nil)
	app.AddRoute("GET", "/user/list/", user.List, nil)
	app.AddRoute("POST", "/user/delete/:id", user.Delete, nil)

	app.Run(":9002")
}
