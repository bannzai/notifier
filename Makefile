APP_NAME=notifier
PORT=5000
CURRENT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

build:
	go build -v -o $(APP_NAME) main.go

clean:
	rm -rf ./$(APP_NAME)

docker: 
	docker build -t $(APP_NAME) .
	docker run -e PORT=$(PORT) -p $(PORT):$(PORT) --rm $(APP_NAME)

heroku: docker
	heroku container:push web

deploy:
	heroku stack:set container
	git push heroku $(CURRENT_BRANCH):master
