package mlog

import (
	"flag"
	"github.com/adayswait/mojo/utils"
	"github.com/golang/glog"
	"os"
)

var log_dir string = "./logs"

func Init() {
	flag.Parse()

	if !utils.IsPathExist(log_dir) {
		err := os.Mkdir(log_dir, os.ModePerm)
		if err != nil {
			glog.Fatal("create log dir failed")
		}
	}
	flag.Lookup("log_dir").Value.Set(log_dir)
	// flag.Lookup("logtostderr").Value.Set("true")
	// flag.Lookup("v").Value.Set("0")
}

var CopyStandardLogTo func(name string) = glog.CopyStandardLogTo
var Error func(args ...interface{}) = glog.Error
var ErrorDepth func(depth int, args ...interface{}) = glog.ErrorDepth
var Errorf func(format string, args ...interface{}) = glog.Errorf
var Errorln func(args ...interface{}) = glog.Errorln
var Exit func(args ...interface{}) = glog.Exit
var ExitDepth func(depth int, args ...interface{}) = glog.ExitDepth
var Exitf func(format string, args ...interface{}) = glog.Exitf
var Exitln func(args ...interface{}) = glog.Exitln
var Fatal func(args ...interface{}) = glog.Fatal
var FatalDepth func(depth int, args ...interface{}) = glog.FatalDepth
var Fatalf func(format string, args ...interface{}) = glog.Fatalf
var Fatalln func(args ...interface{}) = glog.Fatalln
var Flush func() = glog.Flush
var Info func(args ...interface{}) = glog.Info
var InfoDepth func(depth int, args ...interface{}) = glog.InfoDepth
var Infof func(format string, args ...interface{}) = glog.Infof
var Infoln func(args ...interface{}) = glog.Infoln
var Warning func(args ...interface{}) = glog.Warning
var WarningDepth func(depth int, args ...interface{}) = glog.WarningDepth
var Warningf func(format string, args ...interface{}) = glog.Warningf
var Warningln func(args ...interface{}) = glog.Warningln
var Log func(args ...interface{}) = glog.Infoln
