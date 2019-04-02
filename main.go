package main

import (
	"github.com/regod/gwt"
	"github.com/regod/gwt-example/user"
)

func main() {
	app := gwt.New()

	app.AddRoute("/user/create/", user.Create, nil)
	app.AddRoute("/user/update_phone/:id", user.UpdatePhone, nil)
	app.AddRoute("/user/list/", user.List, nil)
	app.AddRoute("/user/delete/:id", user.Delete, nil)

	app.Run(":9002")
}
