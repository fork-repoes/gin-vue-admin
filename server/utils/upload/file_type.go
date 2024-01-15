package upload

import (
	"mime/multipart"
	"net/http"
	"os"
)

func DetermineByFile(file multipart.File) (string, error) {
	// 读取文件前 512 个字节用于判断文件类型
	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// 获取文件的MIME类型
	return http.DetectContentType(buffer), nil
}

func DetermineByPath(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return DetermineByFile(file)
}
