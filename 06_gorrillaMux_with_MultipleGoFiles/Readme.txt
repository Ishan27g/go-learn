http://localhost:80/books/$bookname : POST      - adds $bookname to a map
http://localhost:80/books           : GET       - displays all stored names
http://localhost:80/books/$bookname : DELETE    - deletes $bookname from the map

Importing a custom folder has its rules with capital variable names and import names

With multiple go file, `go build` only verifies if code is compilable. use `go run`