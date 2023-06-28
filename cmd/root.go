/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Toggle bool

var rootCmd = &cobra.Command{
	Use:   "fsm-cobra",
	Short: "Testing FSM",
	Long:  "Interactive testing of FSM",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Is toggle:", Toggle)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&Toggle, "toggle", "t", false, "Help message for toggle")
}
