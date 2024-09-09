package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Println(input string) {
	file, err := os.OpenFile("log/app_log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	writter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(writter)
	log.Println(input)
}

func FatalStr(input string) {
	Println(input)
	os.Exit(1)
}

func FatalErr(err error, exit bool) {
	Println(err.Error())
	if exit {
		panic(err.Error())
	}
}

func Printf(format string, v ...any) {
	var input = fmt.Sprintf(format, v...)
	Println(input)
}
