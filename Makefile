run: sshserver
	./builds/sshserver

sshserver: *.go
	go build -o ./builds