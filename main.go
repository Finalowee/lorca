package main

import (
	`log`
	`os`
	`os/signal`
	`runtime`
	`syscall`

	`github.com/zserge/lorca`

	`supersdk.supersdk.cn/server`
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
	ui, err := lorca.New("", "", GetSystemMetrics(0), GetSystemMetrics(1), args...)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = ui.Close()
	}()

	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: ui.Load("data:text/html," + url.PathEscape(html))

	err = ui.Load(server.WS.Run("F:/golang/supersdk.supersdk.cn/server/www", 8080))
	if err != nil {
		log.Fatal(err)
	}
	// Wait until the interrupt signal arrives or browser window is closed
	sigC := make(chan os.Signal)
	signal.Notify(sigC, os.Interrupt)
	select {
	case <-sigC:
	case <-ui.Done():
	}
	log.Println("exiting...")
}
