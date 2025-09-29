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
-h, --help		|	Show this message
-p, --port		|	Port to run the server on (default : 8080)
-t, --tokenName		|	The name of the token that must be extracted (default : "csrf")
-l, --listChar		|	List of char that can compose the token (default : "abcdefghijklmnopqrstuvwxyz0123456789")
-v, --verbose		|	Activate verbose mode (default : False)
-n, --tokenSize		|	The length of the secret. Should be equal or bigger than the real length of the secret (default : 30)
-s, --prefixToken	|	Specify the characters that prefixed the token (default : "")`,
}

func Execute() {
	err := rootCmd.Execute()
	server.LaunchServer(port)

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&server.TokenLen, "tokenSize", "n", 30, "The size of the secret. Should be equal or bigger than the real size of the secret")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Port to run the webpage on")
	rootCmd.PersistentFlags().StringVarP(&server.TokenName, "tokenName", "t", "csrf", "The name of the secret that must be extracted")
	rootCmd.PersistentFlags().StringVarP(&server.ListOfChar, "listChar", "l", "abcdefghijklmnopqrstuvwxyz0123456789", "List of char that can compose the secret")
	rootCmd.PersistentFlags().BoolP("help", "h", true, "Show the help message")
	rootCmd.PersistentFlags().BoolVarP(&server.Verbose, "verbose", "v", false, "Activate verbose mode")
	rootCmd.PersistentFlags().StringVarP(&server.TokenValue, "prefixToken", "s", "", "Specify the characters that prefixed the token")

}
