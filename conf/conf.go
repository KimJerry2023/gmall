package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode     string
	HttpPort    string
	UploadModel string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	PhotoHost        string
	ProductPhotoPath string
	AvatarPath       string

	AccessKey string
	SecretKey string
	Bucket    string
	Endpoint  string
)

func Init() {
	// 从本地读取环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}
	LoadServer(file)
	LoadMysqlData(file)
	LoadAliyun(file)
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		panic(err)
	}
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	UploadModel = file.Section("service").Key("UploadModel").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadAliyun(file *ini.File) {
	AccessKey = file.Section("aliyun").Key("AccessKey").String()
	SecretKey = file.Section("aliyun").Key("AccessSecret").String()
	Bucket = file.Section("aliyun").Key("Bucket").String()
	Endpoint = file.Section("aliyun").Key("Endpoint").String()
}
