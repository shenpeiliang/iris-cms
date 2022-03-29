package admin

import (
	"cms/util"
	"fmt"
	"mime/multipart"
	"os"
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

	destDirectory := "./" + filepath.Join("static/upload", imgType)

	//目录是否存在
	_, err := os.Stat(destDirectory)
	if os.IsNotExist(err) {
		if e := os.MkdirAll(destDirectory, os.ModePerm); e != nil {
			fmt.Println(e.Error())
		}

	}

	ctx.UploadFormFiles(destDirectory, beforeSave)

	util.Response.Success(ctx, "上传成功", iris.Map{
		"path": "/" + filepath.Join(destDirectory, upload.FileName),
	})

}

//富文本编辑器
func (u Uploadify) Editor(ctx iris.Context) {

	//10M
	ctx.SetMaxRequestBodySize(10 << 20)

	imgType := "editor"

	destDirectory := "./" + filepath.Join("static/upload", imgType)

	//目录是否存在
	_, err := os.Stat(destDirectory)
	if os.IsNotExist(err) {
		if e := os.MkdirAll(destDirectory, os.ModePerm); e != nil {
			fmt.Println(e.Error())
		}

	}

	ctx.UploadFormFiles(destDirectory, beforeSave)

	ctx.JSON(iris.Map{
		"location": "/" + filepath.Join(destDirectory, upload.FileName),
	})

}

func beforeSave(ctx iris.Context, file *multipart.FileHeader) {

	//纳秒
	unixTime := time.Now().UnixNano()

	//file.Header文件类型

	file.Filename = gconv.String(unixTime) + ".jpg"

	upload.FileName = file.Filename
}
