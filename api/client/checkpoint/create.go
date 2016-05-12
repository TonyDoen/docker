package checkpoint

import (
	"golang.org/x/net/context"

	"github.com/docker/docker/api/client"
	"github.com/docker/docker/cli"
	"github.com/docker/engine-api/types"
	"github.com/spf13/cobra"
)

type createOptions struct {
	container    string
	checkpoint   string
	leaveRunning bool
}

func newCreateCommand(dockerCli *client.DockerCli) *cobra.Command {
	var opts createOptions

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a checkpoint from a running container",
		Args:  cli.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.container = args[0]
			opts.checkpoint = args[1]
			return runCreate(dockerCli, opts)
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&opts.leaveRunning, "leave-running", false, "leave the container running after checkpoing")

	return cmd
}

func runCreate(dockerCli *client.DockerCli, opts createOptions) error {
	client := dockerCli.Client()

	checkpointOpts := types.CheckpointCreateOptions{
		CheckpointID: opts.checkpoint,
		Exit:         opts.leaveRunning,
	}

	err := client.CheckpointCreate(context.Background(), opts.container, checkpointOpts)
	if err != nil {
		return err
	}

	return nil
}
