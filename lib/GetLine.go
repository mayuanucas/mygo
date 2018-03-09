package lib

import (
	"bufio"
	"strings"
	"os"
)

func GetLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}
