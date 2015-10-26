./dependencies.sh

rm -rf ./bin

BIN="pgp-smtpd"

for GOOS in windows darwin linux; do
    for GOARCH in 386 amd64; do
      echo "Building $GOOS-$GOARCH"
      export GOOS=$GOOS
      export GOARCH=$GOARCH
      if [ "$GOOS" == "windows" ]
      then
        go build -o bin/${BIN}-$GOOS-$GOARCH.exe
      else
        go build -o bin/${BIN}-$GOOS-$GOARCH
      fi
    done
done
 
