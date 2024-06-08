default:
  just --list

run:
  goht generate && go build && ./QR-Postcard.exe