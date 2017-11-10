.PHONY: run docker-run container docker-stop

run:
	go run index.go

# Parameters
TAG=devel
IMAGE=kube-bot

docker-run: container docker-stop
	docker run -it --rm --name ${IMAGE} ${IMAGE}:${TAG}

container:
	docker build -t ${IMAGE}:${TAG} .

docker-stop:
	bin/docker-stop
