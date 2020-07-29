package steg

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/auyer/steganography"
)

func Encode(imagePath, outputPath, messageFile string) {

	message, err := ioutil.ReadFile("test.txt")
	checkError(err)

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
