package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"time"
)

func openExe(pathStr, arge string) {
	log.Println("OPEN:", pathStr)
	cmd := exec.Command(pathStr, arge)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("command finish...")
}

var pathStr string

func main() {
	pathStr = "C:\\Users\\eyelash500\\AppData\\Local\\LINE\\bin\\LineLauncher.exe"
	openExe(pathStr, "")
	pathStr = "C:\\Program Files (x86)\\Skype\\Phone\\Skype.exe"
	openExe(pathStr, "")
	pathStr = "C:\\Users\\eyelash500\\AppData\\Local\\slack\\app-2.1.2\\slack.exe"
	openExe(pathStr, "")

	log.Printf("To key 'Enter' to leave")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	log.Printf("Bye...")
	time.Sleep(3 * time.Second)

}
