package cli

import (
	"io"
	"os"
)

type Cli interface {
	In() io.Reader
	Out() io.Writer
	Err() io.Writer
}

type DefaultCli struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

func (d *DefaultCli) In() io.Reader {
	return d.stdin
}

func (d *DefaultCli) Out() io.Writer {
	return d.stdout
}

func (d *DefaultCli) Err() io.Writer {
	return d.stderr
}

func NewDefaultCli() *DefaultCli {
	return &DefaultCli{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
}
