package skylog

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"runtime"
	"path/filepath"
)

var (
	errorColor = "\033[31m"
	resetColor = "\033[0m"
)

//SetLogFile function
func SetLogFile(fileName string) (*os.File, error) {
	timeStr := strings.Replace(time.Now().Format("2006-01-02 15:04:05.99"), "-", "_", -1)
	timeStr = strings.Replace(timeStr, ":", "", -1)
	timeStr = strings.Replace(timeStr, " ", "_", -1)
	timeStr = strings.Replace(timeStr, ".", "_", -1)
	f, err := os.OpenFile(fmt.Sprintf("%v_%v.log", fileName, timeStr), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	errorColor = ""
	resetColor = ""
	log.SetOutput(f)
	return f, nil
}

//Info function
func Info(obj ...interface{}) {
	logger(resetColor, "INFO", obj...)
}

func logger(color, section string, obj ...interface{}) {
	log.SetPrefix("[" + section + "]: ")
	_, filename, line, _ := runtime.Caller(2)
	filename = beautyFilename(filename)
	if len(obj) == 2 {
		log.Printf(color+"[%v]: %v(%v): %v\n",  obj[0], filename, line, obj[1])
	} else if len(obj) == 1 {
		log.Printf(color+"%v(%v): %v\n", filename, line, obj[0])
	} else {
		log.Printf(color+"%v(%v): %v\n", filename, line, obj)
	}
	log.Print(resetColor)
}

//Error function
func Error(obj ...interface{}) {
	logger(errorColor, "ERROR", obj...)
}

//Errorf function
func Errorf(format string, obj ...interface{}) {
	log.Printf(errorColor+format, obj...)
}

//Infof function
func Infof(format string, obj ...interface{}) {
	log.Printf(resetColor+format, obj...)
}

//DetailInfo function
func DetailInfo(obj ...interface{}) {
	detialLogger(resetColor, obj...)
}

//DetailError function
func DetailError(obj ...interface{}) {
	detialLogger(errorColor, obj...)
}

func detialLogger(color string, obj ...interface{}) {
	_, filename, line, _ := runtime.Caller(2)
	filename = beautyFilename(filename)
	if len(obj) == 2 {
		log.Printf(color+"%v(%v): [%v]: %#v\n", filename, line, obj[0], obj[1])
	} else if len(obj) == 1 {
		log.Printf(color+"%v(%v): %#v\n", filename, line, obj[0])
	} else {
		log.Printf(color+"%v(%v): %#v\n", filename, line, obj)
	}
	log.Print(resetColor)
}

//Fatal function
func Fatal(obj ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	if len(obj) > 0 && obj[0] != nil {
		if len(obj) == 2 {
			log.Fatal(fmt.Sprintf("%v: (%v): [%v]: %v\n", filename, line, obj[0], obj[1]))
		} else {
			log.Fatal(fmt.Sprintf("%v: (%v) %v\n", filename, line, obj))
		}
	}

}

//ReturnError function
func ReturnError (err error) error {
	Error(err)
	return err
}

func beautyFilename (filename string) string {
	if len(filename) < 1 {
		return filename
	}

	parts := strings.Split(filename, "/")
	if len(parts) < 4 {
		return filename
	}

	return filepath.Join(parts[len(parts)-4], parts[len(parts)-3], parts[len(parts)-2], parts[len(parts)-1])
}