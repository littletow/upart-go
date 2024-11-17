package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mdCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Markdown帮助文档",
	Long:  `Markdown常用语法帮助文档`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("介绍Markdown用法")

	},
}

func init() {
	rootCmd.AddCommand(mdCmd)
}
