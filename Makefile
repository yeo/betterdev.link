GIT_COMMIT := $(shell git rev-list -1 HEAD)
VERSION ?= 0.2
DOCKER_REPO := quay.io/yeospace/betterdev

osx:
	cd cmd && go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT)" -o ../bd

.PHONY: linux
linux:
	cd cmd && GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION) -X main.GitCommit=$(GIT_COMMIT)" -o ../linux

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
	docker build -t ${DOCKER_REPO}:${GIT_COMMIT} .
	docker tag ${DOCKER_REPO}:${GIT_COMMIT} ${DOCKER_REPO}:latest
	docker push ${DOCKER_REPO}:${GIT_COMMIT}
	docker push ${DOCKER_REPO}:latest

k8s-deploy:
	kubectl -n yeo set image deployment/betterdev betterdev=${DOCKER_REPO}:${GIT_COMMIT}

release: linux docker k8s-deploy

fanout:
	export ISSUE
	export PROFILE
	cat k8s/job.yaml | envsubst | kubectl -n yeo delete -f - || true
	cat k8s/job.yaml | envsubst | kubectl -n yeo create -f -

nginx:
	docker build -t ${DOCKER_REPO}:static -f Dockerfile-static .
	docker push ${DOCKER_REPO}:static
