package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:     "open <clusterConfigName>",
		Aliases: []string{"o"},
		Example: "k open clusterConfigName",
		Args:    cobra.MinimumNArgs(1),
		Run:     openK9s,
		Short:   "open k9s by clusterConfigName",
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:     "prepare",
		Aliases: []string{"a"},
		Args:    cobra.MinimumNArgs(1),
		Run:     prepare,
		Short:   "prepare install software",
	})

}

func openK9s(cmd *cobra.Command, args []string) {

	cfName := args[0]

	fmt.Printf("Change k8s default config to [%s]", cfName)

	runCmd("kubecm", "s", args[0])
	fmt.Printf("Opening k9s for [%s]", cfName)
	runCmd("k9s", "-n", "default")

}



func runCmd(cmdName string, params ...string) error {
	cmd := exec.Command(cmdName, params...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return err
	}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	err = cmd.Wait()
	return err

}
