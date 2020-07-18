http://localhost:80/books/$bookname : POST      - adds $bookname to a map
http://localhost:80/books           : GET       - displays all stored names
http://localhost:80/books/$bookname : DELETE    - deletes $bookname from the map

-Importing a custom folder has its rules with capital variable names and import names
-With multiple go files, `go build mux-basic.go` only verifies if code is compilable. use `go run mux-basic.go`
-To build binary, use `go build` which creates binary named as the root folder

#Docker - localhost:81
docker build -t mux_rest_api .
docker run -it -p 81:80 mux_rest_api
