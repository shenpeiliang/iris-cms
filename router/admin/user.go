package admin

import (
	"cms/controller/admin"

	"github.com/kataras/iris/v12"
)

func UserController(party iris.Party) {
	party.PartyFunc("/user", func(user iris.Party) {
		userController := admin.User{}

		user.Get("/login", userController.Login)
		user.Get("/index", userController.Index)
		user.Get("/check", userController.Check)
	})
}
