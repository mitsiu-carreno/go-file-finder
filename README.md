# Go File Finder

Helper to retreive a list of documents from mongo and check if the file exists at a given path

## Deployment

For deployment, build the code into the target architecture
```
$ GOOS=linux GOARCH=amd64 go build main.go
```

For runinig the build, first set the variables required (MAIN_DB_HOST, MAIN_DB_DB, MAIN_DB_USER, MAIN_DB_PASSWORD, MAIN_DB_COLLECTION, FILE_INPUT...)
```
$ source .envfile
$ ./main
```
