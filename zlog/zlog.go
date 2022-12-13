package zlog

import (
	"github.com/golang/glog"
)

func Info(args ...interface{}) {
	glog.Info(args...)
}

func Infoln(args ...interface{}) {
	glog.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func Warning(args ...interface{}) {
	glog.Warning(args...)
}

func Warningln(args ...interface{}) {
	glog.Warningln(args...)
}

func Warningf(format string, args ...interface{}) {
	glog.Warningf(format, args...)
}

func Error(args ...interface{}) {
	glog.Error(args...)
}

func Errorln(args ...interface{}) {
	glog.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	glog.Fatal(args...)
}

func Fatalln(args ...interface{}) {
	glog.Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args...)
}

func Flush() {
	glog.Flush()
}
