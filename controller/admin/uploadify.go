package admin

import (
	"cms/util"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kataras/iris/v12"
)

type Uploadify struct {
	FileName string
}

var upload = new(Uploadify)

//上传文件
func (u Uploadify) Upload(ctx iris.Context) {

	//10M
	ctx.SetMaxRequestBodySize(10 << 20)

	imgType := ctx.Params().GetStringDefault("img_type", "default")

	destDirectory := filepath.Join("./upload", imgType)
	ctx.UploadFormFiles(destDirectory, beforeSave)

	util.Response.Success(ctx, "上传成功", iris.Map{
		"path": filepath.Join(destDirectory, upload.FileName),
	})

}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {

	unixTime := time.Now().Unix()

	file.Filename = gconv.String(unixTime) + "-" + file.Filename

	upload.FileName = file.Filename
}
