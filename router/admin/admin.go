package admin

import (
	"github.com/kataras/iris/v12"
)

func InitAdmin(party iris.Party) {
	party = party.Party("/admin")
	RegisterUser(party)
	RegisterLogin(party)
	RegisterArticle(party)
	RegisterUploadify(party)
}
