package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type MineFormatter struct{}

const TimeFormat = "2006-01-02 15:04:05"

func (s *MineFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	msg := fmt.Sprintf("[%s] [%s] %s\n", time.Now().Local().Format(TimeFormat), strings.ToUpper(entry.Level.String()), entry.Message)

	return []byte(msg), nil
}