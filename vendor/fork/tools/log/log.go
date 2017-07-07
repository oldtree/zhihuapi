package log

import (
	"fmt"
	"runtime"
	"time"
)

func Info(v ...interface{}) {
	funcName, _, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("%s [Info]%s [%d]: ", time.Now().String(), runtime.FuncForPC(funcName).Name(), line)
		fmt.Println(v...)
	}
}

func Debug(v ...interface{}) {
	funcName, _, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("%s [Debug]%s [%d]: ", time.Now().String(), runtime.FuncForPC(funcName).Name(), line)
		fmt.Println(v...)
	}
}

func Warn(v ...interface{}) {
	funcName, _, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("%s [Warning]%s [%d]: ", time.Now().String(), runtime.FuncForPC(funcName).Name(), line)
		fmt.Println(v...)
	}
}

func Critical(v ...interface{}) {
	funcName, _, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("%s [Critic]%s [%d]: ", time.Now().String(), runtime.FuncForPC(funcName).Name(), line)
		fmt.Println(v...)
	}
}

func Error(v ...interface{}) {
	funcName, _, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("%s [Error]%s [%d]: ", time.Now().String(), runtime.FuncForPC(funcName).Name(), line)
		fmt.Println(v...)
	}

}

func TraceAll() {
	fmt.Printf("%s TraceAll: ", time.Now().String())
	for i := 1; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Println(file, line)
	}
}
