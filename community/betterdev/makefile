DOCKER_NAME = bd

build-docker:
	docker build -t $(DOCKER_NAME) .

repl:
	docker run --rm -it -v `pwd`:/app $(DOCKER_NAME) /bin/bash -l
