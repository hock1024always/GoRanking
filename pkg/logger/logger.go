package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime/debug"
	"time"
)

// 初始化日志设置
func init() {
	// 设置日志的 JSON 格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

// 写入程序员自定义的日志
func Write(msg string, filename string) {
	setOutPutFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

// Debug 级别日志
func Debug(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args)
}

// Info 级别日志
func Info(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Info(args)
}

// Warn 级别日志
func Warn(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Warn(args)
}

// Fatal 级别日志
func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Fatal(args)
}

// Error 级别日志
func Error(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Error(args)
}

// Panic 级别日志
func Panic(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.PanicLevel, "panic")
	logrus.WithFields(fields).Panic(args)
}

// Trace 级别日志
func Trace(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.TraceLevel, "trace")
	logrus.WithFields(fields).Trace(args)
}

// 设置日志输出文件 在各个函数方法中调用
func setOutPutFile(level logrus.Level, logName string) {
	// 创建日志目录
	logDir := "./runtime/log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", logDir, err))
		}
	}

	// 获取当前日期字符串
	timeStr := time.Now().Format("2006-01-02")
	fileName := filepath.Join(logDir, logName+"_"+timeStr+".log")

	// 打开日志文件，如果不存在则创建
	var err error
	os.Stderr, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file err:", err)
	}

	// 设置日志输出到文件
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	return
}

// 创建了success开头的文件 在logger中以中间件的形式调用
func LoggerToFile() gin.LoggerConfig {
	logDir := "./runtime/log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", logDir, err))
		}
	}

	// 获取当前日期字符串
	timeStr := time.Now().Format("2006-01-02")
	fileName := path.Join(logDir, "success_"+timeStr+".log")

	os.Stdout, _ = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	var conf = gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.TimeStamp.Format(time.RFC1123),
				param.ClientIP,
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, os.Stderr),
	}
	return conf
}

// 将报错放在报文中返回回来 在logger中以中间件的形式调用
func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if _, errDir := os.Stat("./runtime/log"); os.IsNotExist(errDir) {
				errDir := os.MkdirAll("./runtime/log", 0777)
				if errDir != nil {
					panic(fmt.Errorf("create log dir '%s' error: %s", "./runtime/log", err))
				}
			}
			timeStr := time.Now().Format("2006-01-02")
			//文件名
			fileName := path.Join("./runtime/log", timeStr+".log")

			f, errFile := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if errFile != nil {
				fmt.Println(errFile)
			}
			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			//写入信息
			f.WriteString("panic error tome:" + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v", err) + "\n")
			f.WriteString("stacktrace from panic" + string(debug.Stack()) + "\n")
			f.Close()
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("%v", err),
			})
			//终止后续接口调用，不加入recover到异常之后，还会继续执行接口中的后续代码
			c.Abort()
		}
	}()

	c.Next()
}
