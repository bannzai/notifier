APP_NAME=notifier
DEVPORT=5000
CURRENT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

build:
	go build -v -o $(APP_NAME) main.go

test: 
	go test ./...

clean:
	rm -rf ./$(APP_NAME)

docker: 
	docker build -t $(APP_NAME) .
	docker run -e PORT=$(DEVPORT) -p $(DEVPORT):$(DEVPORT) --rm $(APP_NAME)

heroku:
	heroku stack:set container
	git push heroku $(CURRENT_BRANCH):master
heroku-f:
	heroku stack:set container
	git push heroku $(CURRENT_BRANCH):master --force-with-lease
