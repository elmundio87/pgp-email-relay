./dependencies.sh

rm -rf ./release

BIN="pgp-smtpd"
TAG=`git describe --tags`

for GOOS in windows darwin linux; do
    for GOARCH in 386 amd64; do
      echo "Building $GOOS-$GOARCH"
      export GOOS=$GOOS
      export GOARCH=$GOARCH
      if [ "$GOOS" == "windows" ]
      then
        go build -o release/$GOOS-$GOARCH/${BIN}.exe
      else
        go build -o release/$GOOS-$GOARCH/${BIN}
      fi
      zip -j release/${BIN}-${GOOS}-${GOARCH}-${TAG}.zip release/$GOOS-$GOARCH/* README.md smtp.conf.sample generate_keys.sh
    done
done
 
