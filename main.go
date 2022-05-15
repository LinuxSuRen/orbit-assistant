package main

import (
	"github.com/linuxsuren/orbit-assistant/cmd"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use: "orbit",
	}
	root.AddCommand(cmd.NewCommand())
	if err := root.Execute(); err != nil {
		panic(err)
	}
}
