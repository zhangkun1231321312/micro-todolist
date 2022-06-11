package conf

import (
	"fmt"
	"user/model"

	"strings"

	"gopkg.in/ini.v1"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("conf/conf.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查路径:", err)
		return
	}
	LoadMysqlData(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	fmt.Println(path)
	model.Database(path)
}
func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbName = file.Section("mysql").Key(DbName).String()
	DbHost = file.Section("mysql").Key(DbHost).String()
	DbPassWord = file.Section("mysql").Key(DbPassWord).String()
	DbPort = file.Section("mysql").Key(DbPassWord).String()
	DbUser = file.Section("mysql").Key(DbUser).String()
}
