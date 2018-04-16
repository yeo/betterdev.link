GIT_COMMIT := $(shell git rev-list -1 HEAD)
VERSION ?= 0.1

release:
	cd cmd && go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT)" -o ../bd

linux:
	cd cmd && GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT)" -o ../bd

install:
	cp ./bd ~/bin/bd
	chmod +x ~/bin/bd

server:
	go run cmd/server.go

build:
	bd

upload:
	rsync -rap public/* $(DEPLOY_USER)@$(DEPLOY_PATH)

chmod:
	ssh kurei@betterdev.link "$(DEPLOY_CMD)"

deploy: build upload chmod

docker:
	docker build -t yeospace/betterdev .

