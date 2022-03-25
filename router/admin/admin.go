package admin

import (
	"github.com/kataras/iris/v12"
)

func InitAdmin(party iris.Party) {
	RegisterUser(party)
	RegisterArticle(party)
	RegisterUploadify(party)
}
