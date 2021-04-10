
#install base image -> golang alpine3.13
FROM golang:alpine3.13 as build

#inside docker, create new dir and set it as cwd
WORKDIR /app

#copy contents to docker image
COPY . /app/

#inside docker image, download and install dependencies here
RUN apk update && apk add git && go get -d github.com/gin-gonic/gin

#compile to a binary, `go build` results in binary produced as the root folder, binary = app
RUN go build

#expose it to the host network
EXPOSE 8080

#create new stage
FROM alpine:3.13

#inside this stage, create new dir and set it as cwd
WORKDIR /app

#copy contents to from previous stage to this stage (content is just the final binary and .env file))

COPY --from=build /app/ /app/

EXPOSE 8080

#command to run the binary, `go build` results in binary named as the root folder
CMD ["./ginServer"]
