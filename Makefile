.PHONY: run imports lint fmt \
		docker-run container stop \
		app

# Go src packages
PKGS=bot chat config *.go

run:
	go run main.go

imports:
	goimports -w ${PKGS}

lint:
	golint github.com/Noah-Huppert/kube-bot

fmt:
	gofmt -w ${PKGS}

# Parameters
TAG=devel
IMAGE=kube-bot

docker-run: container stop
	docker run -it --rm --name ${IMAGE} ${IMAGE}:${TAG}

container:
	docker build -t ${IMAGE}:${TAG} .

stop:
	bin/stop

# Rkt
app:
	sudo acbuild script --debug containers/app
