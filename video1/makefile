# build:
# 	go build -o main  .
# run:
# 	./main
docker-build:
	docker build -t my-go-app .
docker-run:
	docker run -p 8080:8080 my-go-app 
all: docker-build docker-run