package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func checkErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	log.Println("Save As?--:")
	filename,err := reader.ReadString('\n')
	filename = strings.TrimSpace(filename)
	checkErr(err)

	// create a file 
	imgFile,err := os.Create(filename+".jpg")
	checkErr(err)
	defer imgFile.Close()

	// collect image from url
	fmt.Println("Whats The Image Url?: ")
	url,err:=reader.ReadString('\n')
	checkErr(err)
	url = strings.TrimSpace(url)
	response,err := http.Get(url)
	checkErr(err)

	defer response.Body.Close()

	// translate response into file
	_,err = io.Copy(imgFile,response.Body)
	checkErr(err)

}
