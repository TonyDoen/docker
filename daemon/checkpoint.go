package daemon

import (
	"fmt"

	"github.com/docker/engine-api/types"
)

// CheckpointCreate checkpoints the process running in a container with CRIU
func (daemon *Daemon) CheckpointCreate(name string, config types.CheckpointCreateOptions) error {
	container, err := daemon.GetContainer(name)
	if err != nil {
		return err
	}

	if !container.IsRunning() {
		return fmt.Errorf("Container %s not running", name)
	}

	err = daemon.containerd.CreateCheckpoint(container.ID, config.CheckpointID, container.CheckpointDir(), config.Exit)
	if err != nil {
		return fmt.Errorf("Cannot checkpoint container %s: %s", name, err)
	}

	daemon.LogContainerEvent(container, "checkpoint")

	if config.Exit {
		daemon.Cleanup(container)
	}

	return nil
}

// CheckpointDelete deletes the specified checkpoint
func (daemon *Daemon) CheckpointDelete(name string, checkpoint string) error {
	container, err := daemon.GetContainer(name)
	if err != nil {
		return err
	}

	err = daemon.containerd.DeleteCheckpoint(container.ID, checkpoint, container.CheckpointDir())
	if err != nil {
		return fmt.Errorf("Cannot delete checkpoint %s: %s", checkpoint, err)
	}

	return nil
}

// CheckpointList deletes the specified checkpoint
func (daemon *Daemon) CheckpointList(name string) ([]types.Checkpoint, error) {
	response := []types.Checkpoint{}

	container, err := daemon.GetContainer(name)
	if err != nil {
		return response, err
	}

	checkpoints, err := daemon.containerd.ListCheckpoints(container.ID, container.CheckpointDir())
	if err != nil {
		return response, fmt.Errorf("Cannot list checkpoints %s: %s", name, err)
	}

	for _, c := range checkpoints.Checkpoints {
		response = append(response, types.Checkpoint{
			Name: c.Name,
		})
	}

	return response, nil
}
