package main

import (
	"cms/util"
	"strings"

	"gorm.io/gen"
)

// generate code
func main() {
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "./",
		/* Mode: gen.WithoutContext|gen.WithDefaultQuery*/
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		/* FieldCoverable: true,*/
		// if you want generate field with unsigned integer type, set FieldSignable true
		/* FieldSignable: true,*/
		//if you want to generate index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
		//if you want to generate type tags from database, set FieldWithTypeTag true
		/* FieldWithTypeTag: true,*/
		//if you need unit tests for query code, set WithUnitTest true
		/* WithUnitTest: true, */
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	// db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(util.DB)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	//g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))

	// apply diy interfaces on structs or table models
	//g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

	//类型映射
	dataMap := map[string]func(detailType string) (dataType string){
		"int": func(detailType string) (dataType string) {
			if strings.HasSuffix(detailType, "unsigned") {
				return "uint"
			}
			return "int"
		},
		"tinyint": func(detailType string) (dataType string) {
			return "byte"
		},
	}
	g.WithDataTypeMap(dataMap)

	//文件名
	g.WithFileNameStrategy(func(tableName string) (fieldName string) {
		return string([]byte(tableName)[3:])
	})

	//生成全部表结构
	g.GenerateAllTable()

	// execute the action of code generation
	g.Execute()
}
