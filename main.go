package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/zserge/lorca"

	"lorca/server"
)

func GetSystemMetrics(nIndex int) int {
	ret, _, _ := syscall.NewLazyDLL(`User32.dll`).NewProc(`GetSystemMetrics`).Call(uintptr(nIndex))
	return int(ret)
}

func main() {
	args := make([]string, 0)
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", GetSystemMetrics(0)/2, GetSystemMetrics(1)/2, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = ui.Close()
	}()

	// 将 server 目录启动为静态服务器
	err = ui.Load(server.WS.Run("./server/www", 8080))
	if err != nil {
		log.Fatal(err)
	}
	// 等待结束
	sigC := make(chan os.Signal)
	signal.Notify(sigC, os.Interrupt)
	select {
		case <-sigC:
		case <-ui.Done():
	}
	log.Println("exiting...")
}
