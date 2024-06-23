default:
  just --list

build:
  goht generate && go build

run_windows: build
  ./QR-Postcard.exe

run_linux: build
  ./QR-Postcard