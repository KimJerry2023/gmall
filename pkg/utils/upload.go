package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"gmall/conf"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"time"
)

// UploadProductToLocalStatic 上传到本地文件中
func UploadProductToLocalStatic(file multipart.File, bossId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(bossId))
	basePath := "." + conf.ProductPhotoPath + "boss" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "boss" + bId + "/" + productName + ".jpg", err
}

// UploadAvatarToLocalStatic 上传头像
func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "user" + bId + "/" + userName + ".jpg", err
}

// DirExistOrNot 判断文件是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

// CreateDir 创建文件加
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 7550)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// OssUpload 封装上传图片到阿里云然后返回状态和图片的url, 单张
func OssUpload(fileHeader *multipart.FileHeader) (string, error) {
	var AccessKey = conf.AccessKey
	var AccessKeySecret = conf.SecretKey
	var Endpoint = conf.Endpoint
	var BucketName = conf.Bucket
	fileHandle, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()
	// file 中没有单独列出扩展名，所以此处需要单独取一次
	fileExt := path.Ext(fileHeader.Filename)
	// 此处重命名文件 取此时的时间戳为MD5为上传OSS的文件名
	data := []byte(time.Now().String())
	md5FileName := fmt.Sprintf("%x", md5.Sum(data))
	// 以年月为文件目录进行分类
	tTime := time.Now().Format("200601")
	// 年月/文件名.扩展名（注意不要再定义的目录前面加/）
	ossFilePath := fmt.Sprintf("%s/%s%s", tTime, md5FileName, fileExt)
	client, err := oss.New(Endpoint, AccessKey, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		return "", err
	}
	err = bucket.PutObject(ossFilePath, fileHandle)
	if err != nil {
		return "", err
	}
	fileRes := fmt.Sprintf("https://%s.%s/%s", BucketName, Endpoint, ossFilePath)
	return fileRes, nil
}
