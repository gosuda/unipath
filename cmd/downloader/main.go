package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Src string `arg:"" name:"src" help:"Source URL to download from."`
	Dst string `arg:"" name:"dst" help:"Destination URL or local path to download to." optional:"" default:"."`
}

func main() {
	ctx := kong.Parse(&CLI)
	defer ctx.Exit(0)

	fmt.Printf("Source URL: %s\n", CLI.Src)
	fmt.Printf("Destination URL: %s\n", CLI.Dst)

	// download(CLI.Src, CLI.Dst)
}
