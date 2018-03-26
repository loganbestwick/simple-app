# Move the application into the correct package for $GOPATH
DEST=$GOPATH/src/github.com/loganbestwick/simple-app
rm -rf $DEST
mkdir -p $DEST
cp -r * $DEST

# Build the server
go build -o application $DEST/cmd/main.go
