package logger

import (
	"github.com/sirupsen/logrus"
	"go.elastic.co/ecslogrus"
)

var (
	l *logrus.Logger
)

func InitLogger() {
	l = logrus.New()
	l.SetFormatter(&ecslogrus.Formatter{})
	l.SetLevel(logrus.InfoLevel)
}

func Infoln(args ...interface{}) {
	l.Infoln(args...)
}

func Fatalln(args ...interface{}) {
	l.Fatalln(args...)
}

func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}

func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}
