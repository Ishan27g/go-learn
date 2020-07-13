#install base image -> golang v1.13-alpine3.11
FROM golang:1.13-alpine3.11 AS build
#inside docker, create new dir and set it as cwd
WORKDIR /app
#copy contents to docker image
COPY . /app/
#inside docker image, compile to a binary
RUN go build main.go

# This 'main' binary file can now be run on >alpine3.11 images.
# We can now create a new stage without golang:1.13, drastically reducing the image size

#create new stage
FROM alpine:3.11
#inside this stage, create new dir and set it as cwd
WORKDIR /app
#copy contents to from previous stage to this stage (content is just the final binary)
COPY --from=build /app/ /app/
#command to run 'main'
CMD ["/app/main"]