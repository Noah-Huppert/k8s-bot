.PHONY: run container stop

run: container stop
	docker run --rm --name k8s-bot k8s-bot:devel

container:
	docker build -t k8s-bot:devel .

stop:
	docker stop k8s-bot || docker rm k8s-bot || true
