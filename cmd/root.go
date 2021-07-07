package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k",
	Short: "k 是一个k8s多集群管理工具",
	Long: `k 借助 kubecm k9s 来管理k8s多集群`,

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}