package TextColor

import (
	"fmt"
)

func Color(colorString string) func(...interface{}) string {
  sprint := func(args ...interface{}) string {
    return fmt.Sprintf(colorString,
      fmt.Sprint(args...))
  }
  return sprint
}

func RED(text string) {
	fmt.Println(Color("\033[1;31m%s\033[0m")(text))
}

func BLACK(text string) {
	fmt.Println(Color("\033[1;30m%s\033[0m")(text))
}

func GREEN(text string) {
	fmt.Println(Color("\033[1;32m%s\033[0m")(text))
}

func YELLOW(text string) {
	fmt.Println(Color("\033[1;33m%s\033[0m")(text))
}

func PURPLE(text string) {
	fmt.Println(Color("\033[1;34m%s\033[0m")(text))
}

func MAGENTA(text string) {
	fmt.Println(Color("\033[1;35m%s\033[0m")(text))
}

func TEAL(text string) {
	fmt.Println(Color("\033[1;36m%s\033[0m")(text))
}

func WHITE(text string) {
	fmt.Println(Color("\033[1;37m%s\033[0m")(text))
}