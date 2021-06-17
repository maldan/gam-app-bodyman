package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-caloryman/internal/app/caloryman"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
