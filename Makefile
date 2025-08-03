IMAGE_NAME := ghcr.io/arch-err/webhook-trigger
TAG := dev

build:
	tailwindcss -i app/src/styles.css -o app/static/css/styles.css
	docker build -t $(IMAGE_NAME):$(TAG) .

build-push:
	tailwindcss -i app/src/styles.css -o app/static/css/styles.css
	docker build -t $(IMAGE_NAME):$(TAG) .
	docker push $(IMAGE_NAME):$(TAG)

run:
	docker run -p 8080:80 $(IMAGE_NAME):$(TAG)

clean:
	docker rmi -f $(IMAGE_NAME):$(TAG) || true

.PHONY: build run push clean
