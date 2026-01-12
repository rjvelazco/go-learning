package tasktracer

import (
	"bufio"
	"fmt"
	"os"
)

func App() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
