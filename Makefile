.PHONY: run container stop

run: container stop
	docker run -it --rm --name k8s-bot k8s-bot:devel

container:
	docker build -t k8s-bot:devel .

stop:
	bin/docker-stop
