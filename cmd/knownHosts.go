package cmd

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/amitizle/sshtools/internal/printer"
	"github.com/spf13/cobra"
)

// knownHostsCmd represents the knownHosts command
var (
	knownHostsCmd = &cobra.Command{
		Use:   "known-hosts",
		Short: "Manipulate known_hosts file",
	}

	knownHostsRmCmd = &cobra.Command{
		Use:   "rm",
		Short: "Remove a line or lines from the known_hosts file",
		Run:   rmKnownHosts,
	}

	knownHostsSortCmd = &cobra.Command{
		Use:   "sort",
		Short: "Sort the known_hosts file alphanumerically",
		Run:   sortKnownHosts,
	}

	fileLine    int
	sortReverse bool
)

// TODO:
// 1) Add a file path together with the default ssh-dir
func init() {
	// add known-hosts command
	sshtoolsCmd.AddCommand(knownHostsCmd)
	// add rm command
	knownHostsRmCmd.Flags().IntVarP(&fileLine, "line", "l", -1, "Line number of the record to delete (starting at 1)") // TODO lines, not line
	knownHostsRmCmd.MarkFlagRequired("line")
	knownHostsCmd.AddCommand(knownHostsRmCmd)

	// add sort command
	knownHostsCmd.AddCommand(knownHostsSortCmd)
	knownHostsSortCmd.Flags().BoolVarP(&sortReverse, "reverse", "r", false, "Sort the file in reverse order")
}

// TODO extract the read/write file logic
func rmKnownHosts(cmd *cobra.Command, args []string) {
	knownHostsFilePath := path.Join(sshDir, "known_hosts") // TODO not a good place for default value
	knownHostsContent, err := ioutil.ReadFile(knownHostsFilePath)
	if err != nil {
		printer.Errorf("Cannot read %s, %v", knownHostsFilePath, err)
		os.Exit(1)
	}

	knownHostsLines := strings.Split(string(knownHostsContent), "\n")
	resultFileContent := make([]string, 0)
	for currLineIndex, currLine := range knownHostsLines {
		if currLineIndex+1 == fileLine {
			printer.Warnf("Removing line:\n%s\n", currLine)
		} else {
			resultFileContent = append(resultFileContent, currLine)
		}
	}

	knownHostsNewContent := strings.Join(resultFileContent, "\n")
	err = ioutil.WriteFile(knownHostsFilePath, []byte(knownHostsNewContent), 0644)
	if err != nil {
		printer.Errorf("Cannot write to %s, %v", knownHostsFilePath, err)
	}
}

// TODO implement
func sortKnownHosts(cmd *cobra.Command, args []string) {
	printer.Infof("ssh dir is %s\n", sshDir)
	printer.Info("Sort called")
}
