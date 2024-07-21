win:
	rm -f game.exe
	GOOS=windows GOARCH=amd64 go build -o bin/collapse.exe main.go
	cp bin/collapse.exe /mnt/c/temp
	rm -f bin/game.exe
clean:
	rm -f bin/game.exe
