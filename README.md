# IceCream File Manager
IceCream File Manager is a basic Go Rest API that lists, views, updates and deletes files in a controlled environment.

## Requirements
- go 1.19+
Get started with Go [Here](https://go.dev/learn/)

## Build and Run
Build the main.go file into an executable called `main`
```
go build cmd/main/main.go
```
Now you can run `main` which will make the API available at `http://localhost:8080`
```
./main
```

## Usage
### List All Files `/files/`
This returns a json response containing the names of all files saved in the set directory.
```
curl http://localhost:8080/files/
[
    {
        "name": "example.txt"
    },
    {
        "name": "example2.txt"
    }
]
```
If the folder is empty, files will return `null`

### Create New File `/files/new/:filename`
This creates a new file in the set directory.
```
curl http://localhost:8080/files/new/filename.txt
```
If the file already exists, an error is logged.

### Save File Lines `/files/:filename/save`
This posts content to a given filename in the set directory.
```
curl -X POST http://localhost:8080/files/filename.txt/save -H 'Content-Type: application/json' -d '[ { "line" : "Some text you want to save here." } ]'
```
If the file does not exist, a new file is created.

### Rename File `/files/replace/:oldfilename/:newfilename`
This changes the name of a given filename.
```
curl http://localhost:8080/files/replace/oldname.txt/newname.txt
```
If the file does not exist, an error is logged.

### Delete File `/files/delete/:filename`
This removed a given filename from the set directory.
```
curl http://localhost:8080/files/delete/filename.txt
```
