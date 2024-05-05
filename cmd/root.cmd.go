package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ev-the-dev/tmplts/cmd/models"
	"github.com/ev-the-dev/tmplts/cmd/utils"
	"github.com/ev-the-dev/tmplts/tui"
	"github.com/spf13/cobra"
)

var version string = "023fjk32208231kfajhdfkj23"

var (
	rootCmd = &cobra.Command{
		Use: "tmplts",
		RunE: func(cmd *cobra.Command, args []string) error {
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
	versionFlag, versionFlagError := cmd.Flags().GetBool("version")
	if versionFlagError != nil {
		return versionFlagError
	}

	if versionFlag {
		printVersion()
		return nil
	}

	printWelcomeMessage()

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
			Prettier:   true,
			Swc:        true,
			Typescript: true,
		}

		writeFiles(&questionnaire)
	} else {
		requestUserInput()
	}

	return nil
}

func printVersion() {
	fmt.Println(version)
}

func init() {
	rootCmd.Flags().BoolP("all", "a", false, "Auto-Generates all config files")
	rootCmd.Flags().BoolP("version", "v", false, "Current version of TmplTS CLI")
}

func printWelcomeMessage() {
	fmt.Println("**********************************************")
	fmt.Println("*                                            *")
	fmt.Println("*   Welcome to TmplTS! Let's get started.    *")
	fmt.Println("*                                            *")
	fmt.Println("**********************************************")
}

func requestUserInput() {
	wizardAnswers := tui.New()
	writeFiles(wizardAnswers)
}

func writeFiles(userInput *models.UserAnswers) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unable to Get Working Directory: %v", err)
	}
	fmt.Printf("\nCWD:::%v", dir)
	// TODO: Check for conflicts -- perhaps another util?
	utils.GenerateAll(userInput, dir)
}
