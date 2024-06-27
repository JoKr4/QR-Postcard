default:
  just --list

build:
  goht generate && go build

run_windows: build
  ./QR-Postcard.exe

build4linux:
  goht generate && GOOS=linux GOARCH=amd64 go build

deploy4linux: build4linux
  cp QR-Postcard /z/go/bin/QR-Postcard/
  cp config.json /z/go/bin/QR-Postcard/
  # cp postcards.json /z/go/bin/QR-Postcard/ #TODO template
  cp -r resources/ /z/go/bin/QR-Postcard/