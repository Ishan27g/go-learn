go build httpFileUpload.go
[ ! -d "./uploadedFiles" ] && echo "created folder for uploaded files" && mkdir uploadedFiles
echo "run ./httpFileUpload and navigate to http://localhost:8082/ on web browser"