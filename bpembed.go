package bpembed

import (
	"bufio"
	"os"
)

func BPEmbed(filename string) string {

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return "NotExists"
		}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
