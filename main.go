package main

import (
	"github.com/alecthomas/kong"
	"github.com/mkadirtan/feed-toolbelt/cmd/find"
)

var cli struct {
	Find find.FindCMD `cmd:""`
}

func main() {
	ctx := kong.Parse(&cli, kong.ShortUsageOnError())
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
