package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if(len(os.Args) < 3) {
		printError();
		return
	}
	
	args := os.Args[1:]
	if args[0] == "-e" || args[0] == "-d" {
		if(args[0] == "-e"){
			fmt.Println(base64.StdEncoding.EncodeToString([]byte(args[1])) )
		} else {
			decodedBytes, err := base64.StdEncoding.DecodeString(args[1])
			if err != nil {
				fmt.Println("Decoding Failed : %v", err)
				return
			}
			fmt.Println(string( decodedBytes ) )
		}
	} else if args[0] == "-f" && args[1] == "-e" && len(args) == 4 {
		inputData, err := ioutil.ReadFile(args[2])
		if err != nil {
			fmt.Println("Encoding Failed : %v", err)
			return
		}
		encodedBytes := base64.StdEncoding.EncodeToString(inputData)
		ioutil.WriteFile(args[3],
			[]byte( encodedBytes ),
			0644)
		fmt.Println("Encoded string written to ", args[3])
	} else if args[0] == "-f" && args[1] == "-d" && len(args) == 4 {
		inputData, err := ioutil.ReadFile(args[2])
		if err != nil {
			fmt.Println("Decoding Failed : %v", err)
			return
		}
		decodedBytes, err := base64.StdEncoding.DecodeString( string(inputData) )
		if err != nil {
			fmt.Println("Decoding Failed : %v", err)
			return
		}
		ioutil.WriteFile(args[3],
			[]byte( decodedBytes ) ,
			0644)
		fmt.Println("Decoded string written to ", args[3])
	} else {
		printError()
	}
	return
}

func printError() {
	fmt.Println("Need minimum 2 args(Order of argument matters)")
	fmt.Println("Example : b64 -e <inputString>")
	fmt.Println("Example : b64 -d <inputEncodedString>")
	fmt.Println("Example : b64 -f -e <inputFileName> <outputFileName>")
	fmt.Println("Example : b64 -f -d <inputFileName> <outputFileName>")
}
