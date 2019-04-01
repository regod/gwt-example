package main

import (
	"gwt"
	"gwt_example/middlewares"
	"gwt_example/user"
)

func main() {
	app := gwt.New()
	app.SetMiddlewares([]gwt.MiddlewareFunc{
		middlewares.MongoDBInit("mongo://127.0.0.1:27017"),
	})
	app.AddRoute("/user/create/", user.Create, nil)
	app.AddRoute("/user/update_phone/:id", user.UpdatePhone, nil)
	app.AddRoute("/user/list/", user.List, nil)
	app.AddRoute("/user/delete/:id", user.Delete, nil)

	app.Run(":9002")
}
