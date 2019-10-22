package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	tk := time.Tick(time.Second * 2)

	for range tk {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

		cmd := exec.Command("go", "env")
		cmd.Stdin = strings.NewReader("some input")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("in all caps: %q\n", out.String())

	}
}
