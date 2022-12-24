# go-dbdb

## Dev
```
git clone git@github.com:yuki777/go-dbdb.git
cd go-dbdb
```

## Build
```
# macos amd64(intel)
GOOS=darwin GOARCH=amd64 go build -o bin/go-dbdb-macos-amd64
# macos arm64(m1,m2,,,)
GOOS=darwin GOARCH=arm64 go build -o bin/go-dbdb-macos-arm64

# linux amd64
GOOS=linux GOARCH=amd64 go build -o bin/go-dbdb-linux-amd64
# linux arm64
GOOS=linux GOARCH=arm64 go build -o bin/go-dbdb-linux-arm64
```
