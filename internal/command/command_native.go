package command

import (
	"context"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"
)

// SignalToProcessGroup is used in tests to disable sending signals to a process group.
var SignalToProcessGroup = true

type NativeCommand struct {
	cfg  *Config
	opts Options

	// cmd is populated when the command is started.
	cmd *exec.Cmd

	cleanFuncs []func() error
}

var _ Command = (*NativeCommand)(nil)

func NewNative(cfg *Config, opts Options) *NativeCommand {
	if opts.Kernel == nil {
		opts.Kernel = NewLocalKernel(nil)
	}

	if opts.Logger == nil {
		opts.Logger = zap.NewNop()
	}

	return &NativeCommand{
		cfg:  cfg,
		opts: opts,
	}
}

func (c *NativeCommand) Running() bool {
	return c.cmd != nil && c.cmd.ProcessState == nil
}

func (c *NativeCommand) Pid() int {
	if c.cmd == nil || c.cmd.Process == nil {
		return 0
	}
	return c.cmd.Process.Pid
}

func (c *NativeCommand) Start(ctx context.Context) (err error) {
	cfg, cleanups, err := normalizeConfig(
		c.cfg,
		newPathNormalizer(c.opts.Kernel),
		modeNormalizer,
		newArgsNormalizer(c.opts.Session, c.opts.Logger),
		newEnvNormalizer(c.opts.Kernel, c.opts.Session.GetEnv),
	)
	if err != nil {
		return
	}

	c.cleanFuncs = append(c.cleanFuncs, cleanups...)

	stdin := c.opts.Stdin

	if f, ok := stdin.(*os.File); ok && f != nil {
		// Duplicate /dev/stdin.
		newStdinFd, err := dup(f.Fd())
		if err != nil {
			return errors.Wrap(err, "failed to dup stdin")
		}
		closeOnExec(newStdinFd)

		// Setting stdin to the non-block mode fails on the simple "read" command.
		// On the other hand, it allows to use SetReadDeadline().
		// It turned out it's not needed, but keeping the code here for now.
		// if err := syscall.SetNonblock(newStdinFd, true); err != nil {
		// 	return nil, errors.Wrap(err, "failed to set new stdin fd in non-blocking mode")
		// }

		stdin = os.NewFile(uintptr(newStdinFd), "")
	}

	c.cmd = exec.CommandContext(
		ctx,
		cfg.ProgramName,
		cfg.Arguments...,
	)
	c.cmd.Dir = cfg.Directory
	c.cmd.Env = cfg.Env
	c.cmd.Stdin = stdin
	c.cmd.Stdout = c.opts.Stdout
	c.cmd.Stderr = c.opts.Stderr

	// Set the process group ID of the program.
	// It is helpful to stop the program and its
	// children.
	// Note that Setsid set in setSysProcAttrCtty()
	// already starts a new process group.
	// Warning: it does not work with interactive programs
	// like "python", hence, it's commented out.
	// setSysProcAttrPgid(c.cmd)

	c.opts.Logger.Info("starting a native command", zap.Any("config", redactConfig(cfg)))

	if err := c.cmd.Start(); err != nil {
		return errors.WithStack(err)
	}

	c.opts.Logger.Info("a native command started")

	return nil
}

func (c *NativeCommand) Signal(sig os.Signal) error {
	c.opts.Logger.Info("stopping the native command with a signal", zap.Stringer("signal", sig))

	if SignalToProcessGroup {
		// Try to terminate the whole process group. If it fails, fall back to stdlib methods.
		err := signalPgid(c.cmd.Process.Pid, sig)
		if err == nil {
			return nil
		}
		c.opts.Logger.Info("failed to terminate process group; trying Process.Signal()", zap.Error(err))
	}

	if err := c.cmd.Process.Signal(sig); err != nil {
		c.opts.Logger.Info("failed to signal process; trying Process.Kill()", zap.Error(err))
		return errors.WithStack(c.cmd.Process.Kill())
	}

	return nil
}

func (c *NativeCommand) Wait() (err error) {
	c.opts.Logger.Info("waiting for the native command to finish")

	defer func() {
		cleanErr := errors.WithStack(c.cleanup())
		c.opts.Logger.Info("cleaned up the native command", zap.Error(cleanErr))
		if err == nil && cleanErr != nil {
			err = cleanErr
		}
	}()

	var stderr []byte

	err = errors.WithStack(c.cmd.Wait())
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			stderr = exitErr.Stderr
		}
	}

	c.opts.Logger.Info("the native command finished", zap.Error(err), zap.ByteString("stderr", stderr))

	return
}

func (c *NativeCommand) cleanup() (err error) {
	for _, fn := range c.cleanFuncs {
		if errFn := fn(); errFn != nil {
			err = multierr.Append(err, errFn)
		}
	}
	return
}
