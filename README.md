# FyScanner 端口扫描器

![](./assets/icon.png)

### 效果:
![](./assets/cut1.png)
![](./assets/cut2.png)

### 安装:
```shell
git clone https://github.com/mazezen/FyScanner.git
go mod tidy
go run .
```

### 打包
OS
```shell
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os darwin -icon ./assets/Icon.png
```
Windows
```shell
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os windows -icon ./assets/Icon.png
```

Linux
```shell
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os linux -icon ./assets/Icon.png
```
