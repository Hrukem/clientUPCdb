package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InputDateFromConsole() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return strings.Trim(text, "\n")
}
