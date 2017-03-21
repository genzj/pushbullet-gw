package main

import (
	"fmt"
	"os"

	"github.com/genzj/pushbullet-gw/cmds"
)

func main() {
	if err := cmds.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
