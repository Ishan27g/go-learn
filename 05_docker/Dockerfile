#install base
FROM golang:1.13-alpine3.11

#which commands to run
RUN mkdir /app

#copy contents from hostdir to docker
ADD . /app

#inside docker, set cwd
WORKDIR /app

#build app inside docker, creates binary called 'main'
RUN go build main.go

#command to run 'main'
CMD ["/app/main"]