package eas_api

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"testing"
)

func TestGen(t *testing.T) {
	// 初始化数据库连接
	dsn := "wikifx:Wikifx2023@tcp(testdb-mysql.fxeyeinterface.com:3306)/cp_test?charset=utf8&collation=utf8_general_ci&parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	// 创建生成器配置
	g := gen.NewGenerator(gen.Config{

		// 关键：设置实体包路径
		ModelPkgPath: "./internal/data/entity",

		OutPath: "",                 // 关键：禁用查询文件生成
		Mode:    gen.WithoutContext, // 可选：不生成带context的方法
	})

	// 使用数据库连接
	g.UseDB(db)
	g.GenerateAllTable()
	// 执行生成
	g.Execute()
}
