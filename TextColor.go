package TextColor

import (
	"fmt"
)

func RED(text string) {
	fmt.Println("\033[31m" + text)
}
