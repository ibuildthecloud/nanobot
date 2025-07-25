//go:build !windows

package supervise

import (
	"context"
	"os/exec"
	"syscall"

	"github.com/nanobot-ai/nanobot/pkg/system"
)

func Cmd(ctx context.Context, command string, args ...string) *exec.Cmd {
	args = append([]string{"_exec", command}, args...)
	cmd := exec.CommandContext(ctx, system.Bin(), args...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	cmd.Cancel = func() error {
		if cmd.Process != nil {
			// Kill the entire process group
			return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		}
		return nil
	}

	return cmd
}
