package steg

import (
	"log"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getOutputPath(path string) string {
	data := strings.Split(path, "/")
	return data[len(data)-1]
}

func GetOutputPath(path string) string {
	outputPath := getOutputPath(path)
	return GenerateName(outputPath)

}

func GenerateName(path string) string {

	data := strings.Split(path, ".")
	newPath := ""

	for i := 0; i < len(data)-1; i++ {
		newPath += data[i]
	}

	return newPath + "_.png"
}
