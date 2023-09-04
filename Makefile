Name = better_debugger

.PHONY:  build linux win cross withoutwindow clean upx install-lint lint callvis-install callvis proxy pprof

debug:
	go build -ldflags "-s -w" -tags debug

release:
	go build -ldflags "-s -w" -tags release
linux:
	set  GOOS=linux
	go build -ldflags "-s -w" -tags release


win:
	set GOOS=win
	go build -ldflags "-s -w" -tags release -o $(Name).exe
	upx -9 $(Name).exe

cross:linux win

withoutwindow:
	go build -ldflags "-s -w -H=windowsgui" -tags release -o $(Name).exe


clean:
	rm -f *.log
	rm -f $(Name)
	rm -f *.exe
	rm -f *.pprof
	rm -f *.txt

# tools-chain for golang
upx:build
	upx -9 $(Name).exe

install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	golangci-lint run

callvis-install:
	go install github.com/ofabry/go-callvis@master

callvis:
	go-callvis main.go

proxy:
	go env -w  GOPROXY=https://goproxy.io,direct

pprof:
	go tool pprof -http=:8080 *.pprof