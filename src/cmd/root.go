package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var file string

func convertToJSON(c *cobra.Command, args []string) error {
	if file == "" {
		fmt.Println("Please, inform some file to be parsed.")
		os.Exit(0)
	}

	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println("error: %f", err)
	}

	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	json_content := []map[string]interface{}{}
	for fileScanner.Scan() {
		dotenv_line := strings.SplitN(fileScanner.Text(), "=", 2)
		json_content = append(
			json_content,
			map[string]interface{}{
				string(dotenv_line[0]): string(dotenv_line[1]),
			},
		)
	}

	json_line, err := json.MarshalIndent(json_content, "", "  ")
	if err != nil {
		fmt.Printf("error: %f", err)
	}
	fmt.Printf("%s\n", json_line)

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "env2json",
	Short: "Command line tool to convert envfiles into JSON.",
	RunE:  convertToJSON,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&file, "from-file", "f", "", "Dotenv file to be converted into JSON")
}
