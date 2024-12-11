package util

import (
	"fmt"
	"log"
	"time"
)

const (
	Reset = "\033[0m"

	Bright = "\033[1m"
	Dim    = "\033[2m"

	Underline = "\033[4m"
	Reverse   = "\033[7m"
	Hidden    = "\033[8m"

	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

func Output(question int, output string) {
	message := fmt.Sprintf("%sQuestion %d output: %s%s", Underline, question, Cyan+Bright+output+Reset, Reset)
	log.Println(message)
}

func TimeTaken(elapsed time.Duration) {
	message := fmt.Sprintf("%sTime taken: %s%s", Underline, Green+Bright+elapsed.String()+Reset, Reset)
	log.Println(message)
}

func Separator() {
	message := fmt.Sprintf("%s%s%s%s", Yellow, Dim, "----------------------------------------", Reset)
	log.Println(message)
}

func MegaSeparator() {
	log.Println("")
	message := fmt.Sprintf("%s%s%s%s", Blue, Bright, "================================================", Reset)
	log.Println(message)
	log.Println("")
}

func Print(format string, v ...any) {
	message := fmt.Sprintf("%s%s%s%s", Magenta, Bright, format, Reset)
	log.Printf(message, v...)
}

func Fatal(format string, v ...any) {
	message := fmt.Sprintf("%s%s%s%s", Red, Bright, format, Reset)
	log.Fatalf(message, v...)
}
