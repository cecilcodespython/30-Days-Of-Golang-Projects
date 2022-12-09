// day 1

// A Golang Program That Downloads Youtube videos

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var pl = fmt.Println

func checkErr(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pl("Video-Link----//----: ")
	url,err :=reader.ReadString('\n')
	checkErr(err)
	output,err := exec.Command("youtube-dl",url).Output()
	checkErr(err)
	pl(output)

}

