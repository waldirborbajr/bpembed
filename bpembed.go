package bpembed

import (
	"io/ioutil"
	"log"
)

func BPEmbed(filename string) (string){

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error opening file %s", err.Error())
	}
	return string(file)
}
