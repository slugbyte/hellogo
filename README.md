# how to setup a go project

```sh

PROJECT_NAME=$1

go mod init github.com/slugbyte/$PROJECT_NAME
mkdir cmd/$PROJECT_NAME internal/util tests bin
touch cmd/$PROJECT_NAME/main.go
echo "package main" >> cmd/$PROJECT_NAME/main.go
```
