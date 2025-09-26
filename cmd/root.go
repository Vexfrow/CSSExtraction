package cmd

import (
	"CSSExtraction/server"
	"os"

	"github.com/spf13/cobra"
)

var (
	port int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "CSSExtraction",
	Short: "A tool to perform data extraction through CSS imports ",
	Long: `
This tool permits to extracts information using the CSS import functionality on website.
This is done using the Sequential Import Chaining (SIC) technique (@d0nutptr)

Usage : 
CSSExtraction [flags]

Flags:
-h, --help		Show this message
-p, --port		Port to run the server on (default 8080)`,
}

func Execute() {
	err := rootCmd.Execute()
	server.LaunchServer(port)

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Port to run the webpage on")
	rootCmd.PersistentFlags().BoolP("help", "h", true, "Show the help message")
}
