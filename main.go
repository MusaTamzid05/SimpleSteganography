package main

import (
	"flag"
	"fmt"
	"log"
	"steganography_util/steg"
)

func Example() {

	steg.Encode("test.png", "test_output.png", "This is a message")
	fmt.Println(steg.Decode("test_output.png"))
}

func main() {

	encodeFlag := flag.Bool("encode", false, "Set to true if want to encode image")
	decodeFlag := flag.Bool("decode", false, "Set to true if want to decode image")
	imagePath := flag.String("image_path", "", "input image path")
	messageFilePath := flag.String("message_file_path", "", "input message path")

	flag.Parse()

	if *encodeFlag == false && *decodeFlag == false {
		log.Fatalln("Usage:go run main.go -encode -image image_path -massage_file  message_file_path")
	}

	if *encodeFlag == true && *decodeFlag == true {
		log.Fatalln("Usage:go run main.go -encode -image image_path -massage_file  message_file_path")
	}

	if *encodeFlag {

		fmt.Println("Encoding")

		if *imagePath == "" || *messageFilePath == "" {
			log.Fatalln("Usage:go run main.go -encode -image image_path -massage_file  message_file_path")
		}

		outputPath := steg.GetOutputPath(*imagePath)
		steg.Encode(*imagePath, outputPath, *messageFilePath)

	} else if *decodeFlag {
		fmt.Println("Decoding")
		fmt.Println(steg.Decode(*imagePath))
	}

}
