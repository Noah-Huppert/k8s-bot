.PHONY: run imports lint fmt docker-run container stop

run:
	go run main.go

imports:
	goimports -w .

lint:
	golint github.com/Noah-Huppert/kube-bot

fmt:
	gofmt -w .

# Parameters
TAG=devel
IMAGE=kube-bot

docker-run: container stop
	docker run -it --rm --name ${IMAGE} ${IMAGE}:${TAG}

container:
	docker build -t ${IMAGE}:${TAG} .

stop:
	bin/stop
