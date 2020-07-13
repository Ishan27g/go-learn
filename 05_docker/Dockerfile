#install base
FROM golang:1.13-alpine3.11 AS build

#which commands to run
RUN mkdir /app

#copy contents from hostdir to docker
ADD . /app

#inside docker, set cwd
WORKDIR /app

#build app inside docker, creates binary called 'main'
RUN go build main.go


FROM alpine:3.12
WORKDIR /app
COPY --from=build /app/ /app/

#command to run 'main'
CMD ["/app/main"]