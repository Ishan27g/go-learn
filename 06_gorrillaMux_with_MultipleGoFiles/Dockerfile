#install base image -> golang v1.13-alpine3.11
FROM golang:1.13-alpine3.11 as build
#inside docker, create new dir and set it as cwd
WORKDIR /app
#copy contents to docker image
COPY . /app/
#inside docker image, download and install gorilla/mux and git
RUN apk update && apk add git && go get -d github.com/gorilla/mux
#compile to a binary, `go build` results in binary named as the root folder
RUN go build
#expose it to the host network to allow internet access for downloading git, mux
EXPOSE 8080

#create new stage
FROM alpine:3.11
#inside this stage, create new dir and set it as cwd
WORKDIR /app
#copy contents to from previous stage to this stage (content is just the final binary)
COPY --from=build /app/ /app/
#command to run the binary, `go build` results in binary named as the root folder
CMD ["./app"]