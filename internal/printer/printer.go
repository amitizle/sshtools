package printer

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func Error(line string) {
	fmt.Println(aurora.Red(line))
}

func Info(line string) {
	fmt.Println(aurora.Green(line))
}

func Warn(line string) {
	fmt.Println(aurora.Cyan(line))
}
