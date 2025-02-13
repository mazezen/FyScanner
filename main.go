package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("FyScanner端口扫描器")
	w.Resize(fyne.NewSize(600, 600))

	ipEntry := widget.NewEntry()
	ipEntry.SetPlaceHolder("请输入IP地址...")

	startPortEntry := widget.NewEntry()
	startPortEntry.SetPlaceHolder("起始端口")

	endPortEntry := widget.NewEntry()
	endPortEntry.SetPlaceHolder("结束端口")

	resultText := widget.NewMultiLineEntry()
	resultText.MultiLine = true
	resultText.SetPlaceHolder("扫描结果将显示在此处...")

	content := container.NewVBox(
		ipEntry,
		startPortEntry,
		endPortEntry,
		widget.NewButton("开始扫描", func() {
			ip := ipEntry.Text
			startPort, err := strconv.Atoi(startPortEntry.Text)
			if err != nil {
				fmt.Println("Invalid start port")
				return
			}
			endPort, err := strconv.Atoi(endPortEntry.Text)
			if err != nil {
				fmt.Println("Invalid end port")
				return
			}
			log.Println("ip address was: ", ip)
			log.Println("start port was: ", startPort)
			log.Println("end port was: ", endPort)

			var wg sync.WaitGroup
			var mu sync.Mutex
			var scanResult strings.Builder

			for port := startPort; port <= endPort; port++ {
				wg.Add(1)
				go func(p int) {
					defer wg.Done()
					address := fmt.Sprintf("%s:%d", ip, p)
					conn, err := net.Dial("tcp", address)
					if err != nil {
						mu.Lock()
						defer mu.Unlock()
						log.Println(fmt.Sprintf("端口: %d: 关闭\n", p))
						scanResult.WriteString(fmt.Sprintf("端口: %d: 关闭\n", p))
					} else {
						conn.Close()
						mu.Lock()
						defer mu.Unlock()
						log.Println(fmt.Sprintf("端口: %d: 开启\n", p))
						scanResult.WriteString(fmt.Sprintf("端口: %d: 开启\n", p))
					}

					// 更新扫描结果到文本框中
					w.Canvas().Refresh(resultText)
				}(port)
			}
			wg.Wait()

			// 将扫描结果显示在文本框中
			resultText.SetText(scanResult.String())
		}))

	// 创建一个滚动容器并设置其高度
	scrollableResult := container.NewScroll(resultText)
	scrollableContainer := container.NewVBox(
		content,
		scrollableResult,
		layout.NewSpacer(),
	)
	scrollableResult.SetMinSize(fyne.NewSize(600, 300))

	// 设置窗口内容并显示
	w.SetContent(scrollableContainer)
	w.ShowAndRun()
}
