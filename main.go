package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/soft/config"
	"github.com/charmbracelet/soft/server"
	"github.com/spf13/cobra"
)

var (
	Version   = ""
	CommitSHA = ""

	rootCmd = &cobra.Command{
		Use:    "",
		Hidden: false,
		Short:  "A tasty, self-hostable Git server for the command line.",
		Args:   cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := config.DefaultConfig()
			s := server.NewServer(cfg)
			log.Printf("Starting SSH server on %s:%d\n", cfg.Host, cfg.Port)
			return s.Start()
		},
	}
)

func init() {
	if len(CommitSHA) >= 7 {
		vt := rootCmd.VersionTemplate()
		rootCmd.SetVersionTemplate(vt[:len(vt)-1] + " (" + CommitSHA[0:7] + ")\n")
	}
	if Version == "" {
		Version = "unknown (built from source)"
	}
	rootCmd.Version = Version
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
