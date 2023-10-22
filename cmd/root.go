/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/anmol1vw13/word_count/domain"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "word_count",
	Run: func(cmd *cobra.Command, args []string) {

		c := domain.Counter{
			Flag:  &flagSet,
			Files: args,
		}

		c.Start()
		// countFiles, err := c.Count()
		// for _, res := range countFiles {

		// 	if res.Line != 0 {
		// 		fmt.Printf("%d\t", res.Line)
		// 	}
		// 	if res.Word != 0 {
		// 		fmt.Printf("%d\t", res.Word)
		// 	}
		// 	if res.Char != 0 {
		// 		fmt.Printf("%d\t", res.Char)
		// 	}
		// 	fmt.Printf("%s\n", res.Identifier)
		// }
		// if err != nil {
		// 	fmt.Println(err)
		// }

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var flagSet domain.FlagSet

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.word_count.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.SetUsageFunc(nil)
	rootCmd.SetUsageTemplate(
		`WC(1)			    General Commands Manual			 WC(1)

		NAME
			 wc – word, line, character, and byte count
		
		SYNOPSIS
			 wc [-clmw] [file ...]
		
		DESCRIPTION
			 The wc utility displays the number of lines, words, and bytes contained
			 in each input file, or standard input (if no file is specified) to the
			 standard output.  A line is defined as a string of characters delimited
			 by a ⟨newline⟩ character.	Characters beyond the final ⟨newline⟩
			 character will not be included in the line count.
		
			 A word is defined as a string of characters delimited by white space
			 characters.  White space characters are the set of characters for which
			 the iswspace(3) function returns true.  If more than one input file is
			 specified, a line of cumulative counts for all the files is displayed on
			 a separate line after the output for the last file.
		
			 The following options are available:
		
			 -c      The number of bytes in each input file is written to the standard
				 output.  This will cancel out any prior usage of the -m option.
		
			 -l      The number of lines in each input file is written to the standard
				 output.
		
			 -m      The number of characters in each input file is written to the
				 standard output.  If the current locale does not support
				 multibyte characters, this is equivalent to the -c option.  This
				 will cancel out any prior usage of the -c option.
		
			 -w      The number of words in each input file is written to the standard
				 output.
		
			 When an option is specified, wc only reports the information requested by
			 that option.  The order of output always takes the form of line, word,
			 byte, and file name.  The default action is equivalent to specifying the
			 -c, -l and -w options.
		
			 If no files are specified, the standard input is used and no file name is
			 displayed.  The prompt will accept input until receiving EOF, or [^D] in
			 most environments.
		
		ENVIRONMENT
			 The LANG, LC_ALL and LC_CTYPE environment variables affect the execution
			 of wc as described in environ(7).
		
		EXIT STATUS
			 The wc utility exits 0 on success, and >0 if an error occurs.
		
		EXAMPLES
			 Count the number of characters, words and lines in each of the files
			 report1 and report2 as well as the totals for both:
		
			   wc -mlw report1 report2
		
		COMPATIBILITY
			 Historically, the wc utility was documented to define a word as a
			 maximal string of characters delimited by <space>, <tab> or <newline>
			 characters''.  The implementation, however, did not handle non-printing
			 characters correctly so that “  ^D^E  ” counted as 6 spaces, while
			 “foo^D^Ebar” counted as 8 characters.  4BSD systems after 4.3BSD modified
			 the implementation to be consistent with the documentation.  This
			 implementation defines a word'' in terms of the iswspace(3) function,
			 as required by IEEE Std 1003.2 (“POSIX.2”).
		
		SEE ALSO
			 iswspace(3)
		
		STANDARDS
			 The wc utility conforms to IEEE Std 1003.1-2001 (“POSIX.1”).
		
		HISTORY
			 A wc command appeared in Version 1 AT&T UNIX.
		
		macOS 13.5		       February 23, 2005		    macOS 13.5`)

	rootCmd.Flags().BoolVarP(&flagSet.Word, "word", "w", false, "--word")
	rootCmd.Flags().BoolVarP(&flagSet.Char, "character", "c", false, "--word")
	rootCmd.Flags().BoolVarP(&flagSet.Line, "line", "l", false, "--word")
}
