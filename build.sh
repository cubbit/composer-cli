go build -o build/cubbit github.com/cubbit/cubbit/client/cli
env GOOS=windows GOARCH=386 go build -o build/cubbit_386.exe github.com/cubbit/cubbit/client/cli
env GOOS=windows GOARCH=amd64 go build -o build/cubbit_amd64.exe github.com/cubbit/cubbit/client/cli
env GOOS=darwin GOARCH=amd64 go build -o build/cubbit_darwin_amd64 github.com/cubbit/cubbit/client/cli
