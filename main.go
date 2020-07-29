package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/auyer/steganography"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Encode(imagePath, outputPath, message string) {

	inFile, err := os.Open(imagePath)
	checkError(err)

	defer inFile.Close()

	reader := bufio.NewReader(inFile)

	img, err := png.Decode(reader)
	checkError(err)

	writer := new(bytes.Buffer)

	err = steganography.Encode(writer, img, []byte(message))
	checkError(err)

	outFile, err := os.Create(outputPath)
	checkError(err)

	defer outFile.Close()

	writer.WriteTo(outFile)

	fmt.Printf("[*] Message writing to %s\n", outputPath)

}

func Decode(imagePath string) {

	inFile, err := os.Open(imagePath)
	checkError(err)
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, err := png.Decode(reader)
	checkError(err)

	sizeOfMessage := steganography.GetMessageSizeFromImage(img)
	message := steganography.Decode(sizeOfMessage, img)

	fmt.Println(string(message))

}

func Example() {

	Encode("test.png", "test_output.png", "This is a message")
	Decode("test_output.png")
}

func main() {
}
