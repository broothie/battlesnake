
all: build push

build:
	docker build -t gcr.io/battlesnake-258923/battlesnake .

push:
	docker push gcr.io/battlesnake-258923/battlesnake
