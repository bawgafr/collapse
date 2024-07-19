win:
	rm -f game.exe
	GOOS=windows GOARCH=amd64 go build -o bin/collapse.exe main.go
	cp bin/collapse.exe /mnt/c/temp
clean:
	rm -f bin/game.exe
