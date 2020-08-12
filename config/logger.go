package config

import (
	config "IrisFramework/config/Log"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

const LogPath = "./storage/logs"
const FileSuffix = ".log"

var Log = logrus.New()

func InitLog() {
	Log.Out = os.Stdout
	var loglevel logrus.Level
	err := loglevel.UnmarshalText([]byte("info"))
	if err != nil {
		Log.Panicf("设置log级别失败：%v", err)
	}

	Log.SetLevel(loglevel)

	NewSimpleLogger(Log, LogPath, 8)
}

/**
  文件日志
*/
func NewSimpleLogger(log *logrus.Logger, logPath string, save uint) {

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer(logPath, "debug", save), // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer(logPath, "info", save),
		logrus.WarnLevel:  writer(logPath, "warn", save),
		logrus.ErrorLevel: writer(logPath, "error", save),
		logrus.FatalLevel: writer(logPath, "fatal", save),
		logrus.PanicLevel: writer(logPath, "panic", save),
	}, &config.MineFormatter{})

	log.AddHook(lfHook)
}

/**
文件设置
*/
func writer(logPath string, level string, save uint) *rotatelogs.RotateLogs {
	logFullPath := path.Join(logPath, level)
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fileSuffix := time.Now().In(cstSh).Format("2006-01-02") + FileSuffix

	logier, err := rotatelogs.New(
		logFullPath+"-"+fileSuffix,
		rotatelogs.WithLinkName(logFullPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(int(save)),   // 文件最大保存份数
		rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
	)

	if err != nil {
		panic(err)
	}
	return logier
}
