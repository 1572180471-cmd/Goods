/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// pbCmd represents the pb command
var pbCmd = &cobra.Command{
	Use:   "pb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		protocCmd := exec.Command(
			"protoc",
			"--go_out=.",
			"--go_opt=paths=source_relative",
			"--go-grpc_out=.",
			"--go-grpc_opt=paths=source_relative",
			"cmd/goods.proto",
		)
		protocCmd.Stdout = os.Stdout
		protocCmd.Stderr = os.Stderr

		if err := protocCmd.Run(); err != nil {

			cobra.CheckErr(fmt.Errorf("生成PB文件失败: %v", err))
		}
		// 成功消息用cmd.Println输出，而非cobra.CheckErr
		cmd.Println("PB文件生成成功！")
	},
}

func init() {
	rootCmd.AddCommand(pbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
