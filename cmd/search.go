// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-cli/internal/search"
	"os"
	"os/exec"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples and usage of using your command. For example: Cobra is a CLI library for Go that empowers applications. This application is a tool to generate the needed files to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go cli invoke")

		var (
			f     = "temp.md"
			mdSTR = ""
		)

		// 处理掘金 内容
		mdSTR = search.Juejin{}.Run(args[0])
		mdSTR += search.Zhihu{}.Run(args[0])
		mdSTR += search.Csdn{}.Run(args[0])
		mdSTR += search.JianShu{}.Run(args[0])

		os.Create(f)
		file, _ := os.OpenFile(f, os.O_WRONLY|os.O_APPEND, 0666)
		file.WriteString(mdSTR)

		fmt.Println("Generate Successful")
		_, runErr := exec.Command("open", f).Output()
		if runErr != nil {
			panic(runErr)
		}

	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
