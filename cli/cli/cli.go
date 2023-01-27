package cli

import (
	"io"
	"os"
)

type Cli interface {
	In() io.Reader
	Out() io.Writer
	Err() io.Writer

	LogLevel() string
}

type DefaultCli struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer

	loglevel string
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

func (cli *DefaultCli) LogLevel() string {
	return cli.loglevel
}

type FavCli struct {
	DefaultCli

	appMode   string
	configDir string
}

func (cli *FavCli) Apply(opts ...cliOption) {
	for _, opt := range opts {
		opt(cli)
	}
}

func NewFavCli(opts ...cliOption) *FavCli {
	cli := &FavCli{
		DefaultCli: DefaultCli{
			stdin:  os.Stdin,
			stdout: os.Stdout,
			stderr: os.Stderr,
		},
	}

	defaultOptions := []cliOption{
		SetApplicationMode,
		SetLogLevel,
	}

	opts = append(opts, defaultOptions...)
	cli.Apply(opts...)

	return cli
}
