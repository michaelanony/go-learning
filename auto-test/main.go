package main

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

var wg sync.WaitGroup

func getCommand(commandType string, command string) (ret string, err error) {
	cmd := exec.Command("bash", "-c", command)
	//cmd :=exec.Command("cmd","/C",command)
	if commandType == "realtime" {
		stdout, err := cmd.StdoutPipe()
		cmd.Stderr = cmd.Stdout
		if err != nil {
			panic(err)
		}
		if err = cmd.Start(); err != nil {
			return "", err
		}
		for {
			tmp := make([]byte, 1024)
			_, err := stdout.Read(tmp)
			fmt.Print(string(tmp))
			if err != nil {
				break
			}
		}
		if err = cmd.Wait(); err != nil {
			return "", err
		}
		return "", nil
	} else {
		output, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}
		return string(output), nil
	}
}

func getUi() {
	ret, err := getCommand("now", "adb shell cat /sdcard/ui.xml | grep '精选'")
	if err != nil {
		panic(err)
	}
	if ret != "" {

	}
}
func saveUi(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			break
		default:
			time.Sleep(time.Second * 2)
			_, err := getCommand("now", "adb shell uiautomator dump /sdcard/ui.xml")
			if err != nil {
				panic(err)
			}
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go saveUi(ctx)
	go
}
