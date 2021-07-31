GOARCH=arm64 GOOS=linux go build -ldflags "-s -w" -o app .
cp app /mnt/orangepi/root/.gam-app/maldan-gam-app-bodyman-v0.1.7