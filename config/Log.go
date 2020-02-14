package config

import (
	"log"
	"os"
	"time"
)

const DateFormat = "2006-01-02"
const FileType = "log"
const FileDir = "./storage/logs/"
const FilePrefix = "iris-"

/**
文件名
 */
func todayFileName() string {
	return FilePrefix + time.Now().Format(DateFormat) + "." + FileType
}

/**
  文件格式设置
 */
func logFile() *os.File {
	filePath := FileDir + todayFileName()
	// 打开以当前日期为文件名的文件（不存在则创建文件，存在则追加内容）
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	return f
}

func init() {
	// todo 设置日志
	log.SetOutput(logFile())
}
