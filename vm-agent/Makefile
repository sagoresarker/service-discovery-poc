build:
	go build -o agent ./cmd/agent

copy:
	scp ./agent root@10.62.1.231:/root/

copy-env:
	scp ./.env root@10.62.1.231:/root/