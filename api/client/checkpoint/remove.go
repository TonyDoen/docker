package checkpoint

import (
	"golang.org/x/net/context"

	"github.com/docker/docker/api/client"
	"github.com/docker/docker/cli"
	"github.com/spf13/cobra"
)

func newRemoveCommand(dockerCli *client.DockerCli) *cobra.Command {
	return &cobra.Command{
		Use:     "rm CONTAINER CHECKPOINT",
		Aliases: []string{"remove"},
		Short:   "Remove a checkpoint",
		Args:    cli.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRemove(dockerCli, args[0], args[1])
		},
	}
}

func runRemove(dockerCli *client.DockerCli, container string, checkpoint string) error {
	client := dockerCli.Client()
	return client.CheckpointDelete(context.Background(), container, checkpoint)
}
