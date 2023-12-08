package cmd

import (
	"fmt"
	"os"

	"github.com/ehutchllew/template.ts/cmd/models"
	"github.com/ehutchllew/template.ts/cmd/utils"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "tmplts",
		Run: func(cmd *cobra.Command, args []string) {
			printWelcomeMessage()
			userAnswers := &models.UserAnswers{}
			requestUserInput(userAnswers)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("all", "a", false, "tmplts -a")
}

func printWelcomeMessage() {
	fmt.Println("**********************************************")
	fmt.Println("*                                            *")
	fmt.Println("* Welcome to Template.TS! Let's get started. *")
	fmt.Println("*                                            *")
	fmt.Println("**********************************************")
}

func requestUserInput(questionnaire *models.UserAnswers) {
	/*
	 * To be filled out with custom functionality in the future
	 */
	dir, _ := os.Getwd()
	fmt.Printf("\nCWD:::%v", dir)
	// userInput := ""
	// writeFiles(userInput, dir)
}

func writeFiles(userInput string, cwd string) {
	// TODO: Will contain the actual text to write instead of hard coding it inside GenerateRoot
	answers := models.UserAnswers{
		AppName:    "testingggg",
		EsLint:     true,
		Jest:       true,
		Swc:        true,
		Typescript: true,
	}
	utils.GenerateAll(&answers, cwd)
}
