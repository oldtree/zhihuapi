package log

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

var MaxFileSize uint64 = 1024 * 1024 * 1024
var LogStoreDir = "./"
var OutPut io.Writer
var mu sync.Mutex

func init() {
	mu.Lock()
	defer mu.Unlock()
	OutPut = os.Stdout
	fmt.Println("log operation init ok")
}

func createLogDir(dirPath string) error {
	if dirPath == "" {
		dirPath = LogStoreDir
	}
	fi, _ := os.Stat(dirPath)

	if fi != nil {
		if fi.IsDir() {
			return nil
		}
	}
	err := os.MkdirAll(dirPath, 6666)
	if err != nil {
		if !strings.Contains(err.Error(), "exist") {
			return err
		}
	}
	return nil
}

var (
	Pid         = os.Getpid()
	Host, _     = os.Hostname()
	UserRole, _ = user.Current()
	InitOnce    sync.Once
)

func logFileName(ext string, ti time.Time) string {
	name := fmt.Sprintf("%s-%s-%05d-%04d-%02d-%02d-%02d-%02d-%02d",
		Host, UserRole.Username, Pid, ti.Year(), ti.Month(), ti.Day(), ti.Hour(), ti.Minute(), ti.Second())
	return name + "." + ext
}

func Create(ext string, ti time.Time, dir string) (fi *os.File, filename string, err error) {
	defer func() {
		if re := recover(); re != nil {
			fi = nil
			err = fmt.Errorf("create file %s failed ", filename)
		}
	}()
	filename = logFileName(ext, ti)
	err = createLogDir(dir)
	if err != nil {
		return
	}
	fpath := filepath.Join(dir + "/" + filename)
	fi, err = os.Create(fpath)
	if err != nil {
		return nil, filename, err
	}
	mu.Lock()
	defer mu.Unlock()
	OutPut = fi
	return
}

func Info(v ...interface{}) {
	funcName, filename, line, _ := runtime.Caller(1)
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(OutPut, "%s %s [Info] %s [%d]: %v", filename, time.Now().String(), runtime.FuncForPC(funcName).Name(), line, v)
}

func Debug(v ...interface{}) {
	funcName, filename, line, _ := runtime.Caller(1)
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(OutPut, "%s %s [Debug] %s [%d]: %v", filename, time.Now().String(), runtime.FuncForPC(funcName).Name(), line, v)
}

func Warn(v ...interface{}) {
	funcName, filename, line, _ := runtime.Caller(1)
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(OutPut, "%s %s [Warning] %s [%d]: %v", filename, time.Now().String(), runtime.FuncForPC(funcName).Name(), line, v)
}

func Critical(v ...interface{}) {
	funcName, filename, line, _ := runtime.Caller(1)
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(OutPut, "%s %s [Critic] %s [%d]: %v", filename, time.Now().String(), runtime.FuncForPC(funcName).Name(), line, v)
}

func Error(v ...interface{}) {
	funcName, filename, line, _ := runtime.Caller(1)
	fmt.Fprintf(OutPut, "%s %s [Error] %s [%d]: %v", filename, time.Now().String(), runtime.FuncForPC(funcName).Name(), line, v)
}

func TraceAll() {
	mu.Lock()
	defer mu.Unlock()
	for i := 1; ; i++ {
		funcName, filename, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(OutPut, "%s %s [Error] %s [%d]: %v", filename, time.Now().String(), runtime.FuncForPC(funcName).Name(), line)
	}
}
