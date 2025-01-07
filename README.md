## Golang

### Presentation link

https://docs.google.com/presentation/d/1WVvsbvgHKBrNrKtnT4XWRfrsfkNlbw5u6L9O1DeVBn0/edit?usp=sharing


### keywords

break,case ,const,continue,default,else,fallthrough,for, func, if, import, package,return,switch, var  (15 out of 25)

### builtin 
print, println ,complex, real, imag (5 out 18)

### go mod

```bash
go mod init demo
```

### go run 

```bash
go run .
go run main.go
go run main.go anothermain.go
```

### go build 

```bash
go build -o build/hello .
go build main.go anothermain.go
go build -o hello main.go
```

### go build for release

```bash
go build -ldflags="-s -w" -o build/hello-small .   
```

### go compilation targets

```bash
go tool dist list 
```

```bash
GOOS=linux GOARCH=amd64 go build -o build/hello_linux_amd64 .
```

### Go env 

```bash
GOOS
GOARCH
GOROOT
```