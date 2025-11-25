/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runAdd,
}

func runAdd(cmd *cobra.Command, args []string) {
	addTask, err := cmd.Flags().GetBool("create-task")

	if err != nil {
		fmt.Println("error with addTask", err)
	}

	addList, err := cmd.Flags().GetBool("create-list")

	if err != nil {
		fmt.Println("error with addList", err)
	}

	switch {
	case addTask && addList:
		{
			fmt.Println("you cannot add list and task at the same time")
			os.Exit(1)
		}
	case addTask:
		{
			fmt.Println("Adding Task...")
			fmt.Println(args)
		}
	case addList:
		{
			fmt.Println("Adding List...")
			fmt.Println(args)
		}
	}

}

func init() {
	addCmd.Flags().IntP("priority", "p", 2, "Priority of the task")
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("create-list", "y", false, "Create a list")
	addCmd.Flags().BoolP("create-task", "t", false, "Create a task")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
