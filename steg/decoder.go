package steg

import (
	"bufio"
	"image/png"
	"os"

	"github.com/auyer/steganography"
)

func Decode(imagePath string) string {

	inFile, err := os.Open(imagePath)
	checkError(err)
	defer inFile.Close()

	reader := bufio.NewReader(inFile)
	img, err := png.Decode(reader)
	checkError(err)

	sizeOfMessage := steganography.GetMessageSizeFromImage(img)
	message := steganography.Decode(sizeOfMessage, img)

	return string(message)

}
