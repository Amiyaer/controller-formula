package cmd

import (
    "flag"
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"
    "k8s.io/klog/v2"

    "github.com/Amiyaer/controller-formula/pkg/config"
    "github.com/Amiyaer/controller-formula/pkg/controller"
    "github.com/Amiyaer/controller-formula/pkg/version"
)

func main() {
	_ = flag.CommandLine.Parse([]string{})
	var cfg = config.Default()

	cmd := &cobra.Command{
		Use:   "controller-formula",
		Short: "",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			run(cfg)
		},
	}
	
	err := cmd.Execute()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
}

func run(cfg *config.Config) {
	klog.Infof("Version :- %s - commit :- %s", version.GetVersion(), version.GetGitCommit())

	// set your config ang build your controller object
	
	// new your controller and catch error
	err := controller.New(cfg).Run()
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}