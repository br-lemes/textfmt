package cmd

import (
	"bufio"
	"errors"
	"strings"

	"github.com/br-lemes/textfmt/internal/version"
	"github.com/spf13/cobra"
)

var (
	chars bool
	lower bool
	upper bool
	words bool
)

var rootCmd = &cobra.Command{
	Use:   "textfmt",
	Short: "A simple CLI tool for formatting and analyzing text",
	Long: `A simple CLI tool for formatting and analyzing text.

You can provide text as arguments or pipe it via stdin.
Multiple formatting options can be applied simultaneously.`,
	Args: cobra.ArbitraryArgs,
	RunE: runTextfmt,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Version = version.GetVersion()
	rootCmd.Flags().
		BoolVar(&chars, "chars", false, "Count characters in text")
	rootCmd.Flags().
		BoolVarP(&lower, "lower", "l", false, "Convert text to lowercase")
	rootCmd.Flags().
		BoolVarP(&upper, "upper", "u", false, "Convert text to uppercase")
	rootCmd.Flags().
		BoolVar(&words, "words", false, "Count words in text")
}

func runTextfmt(cmd *cobra.Command, args []string) error {
	var text string

	if len(args) > 0 {
		text = strings.Join(args, " ")
	} else {
		scanner := bufio.NewScanner(cmd.InOrStdin())
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		text = strings.Join(lines, "\n")
	}

	if strings.TrimSpace(text) == "" {
		return errors.New("no input text provided")
	}

	if lower {
		text = strings.ToLower(text)
	}
	if upper {
		text = strings.ToUpper(text)
	}

	cmd.Println(text)

	if chars {
		cmd.Printf("Character count: %d\n", len(text))
	}
	if words {
		words := strings.Fields(text)
		cmd.Printf("Word count: %d\n", len(words))
	}
	return nil
}
