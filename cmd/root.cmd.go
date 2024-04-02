package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ehutchllew/template.ts/cmd/models"
	"github.com/ehutchllew/template.ts/cmd/utils"
	"github.com/ehutchllew/template.ts/tui"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "tmplts",
		RunE: func(cmd *cobra.Command, args []string) error {
			printWelcomeMessage()
			checkForFlags(cmd)

			return nil
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func checkForFlags(cmd *cobra.Command) error {
	allFlag, allFlagErr := cmd.Flags().GetBool("all")
	if allFlagErr != nil {
		return allFlagErr
	}

	if allFlag {
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Unable to Get Working Directory: %v", err)
		}
		defaultAppName := path.Base(dir)

		questionnaire := models.UserAnswers{
			AppName:    defaultAppName,
			EsLint:     true,
			Jest:       true,
			Swc:        true,
			Typescript: true,
		}

		writeFiles(&questionnaire)
	} else {
		requestUserInput()
	}

	return nil
}

func init() {
	rootCmd.Flags().BoolP("all", "a", false, "Auto-Generates all config files")
}

func printWelcomeMessage() {
	fmt.Println("**********************************************")
	fmt.Println("*                                            *")
	fmt.Println("* Welcome to Template.TS! Let's get started. *")
	fmt.Println("*                                            *")
	fmt.Println("**********************************************")
}

func requestUserInput() {
	wizardAnswers := tui.New()
	fmt.Printf("WIZARD ANSWERS::: \n%+v", wizardAnswers)
	writeFiles(wizardAnswers)
}

func writeFiles(userInput *models.UserAnswers) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to Get Working Directory: %v", err)
	}
	fmt.Printf("\nCWD:::%v", dir)
	utils.GenerateAll(userInput, dir)
}
