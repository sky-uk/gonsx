package main

import (
	"flag"
	"fmt"
	"github.com/tadaweb/gonsx"
	"os"
)

// ExecFunc executes the function for cli.
type ExecFunc func(client *gonsx.NSXClient, flagSet *flag.FlagSet)

// Command struct - defines a cli command with flags and exec
type Command struct {
	flagSet *flag.FlagSet
	exec    ExecFunc
}

var (
	nsxServer     string
	nsxServerPort int
	nsxUsername   string
	nsxPassword   string
	debug         bool

	commandMap = make(map[string]Command, 0)
)

// RegisterCliCommand - allows additional cli commands to be registered.
func RegisterCliCommand(name string, flagSet *flag.FlagSet, exec ExecFunc) {
	commandMap[name] = Command{flagSet, exec}
}

// InitFlags - initiall cli flags.
func InitFlags() {
	flag.StringVar(&nsxServer, "server", os.Getenv("NSX_SERVER"),
		"NSX API server hostname or address. (Env: NSX_SERVER)")
	flag.IntVar(&nsxServerPort, "port", 443,
		"NSX API server port. Default:443")
	flag.StringVar(&nsxUsername, "username", os.Getenv("NSX_USERNAME"),
		"Authentication username (Env: NSX_USERNAME)")
	flag.StringVar(&nsxPassword, "password", os.Getenv("NSX_PASSWORD"),
		"Authentication password (Env: NSX_PASSWORD)")
	flag.BoolVar(&debug, "debug", false, "Debug output. Default:false")
}

func usage() {
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "  Commands:\n")
	for name := range commandMap {
		fmt.Fprintf(os.Stderr, "    %s\n", name)
	}
}

func main() {
	InitFlags()
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		os.Exit(2)
	}

	command := flag.Arg(0)
	cmd, inMap := commandMap[command]
	if !inMap {
		usage()
		os.Exit(2)
	}

	flagSet := cmd.flagSet
	flagSet.Parse(flag.Args()[1:])

	client := gonsx.NewNSXClient(nsxServer, nsxUsername, nsxPassword, true, debug)

	cmd.exec(client, flagSet)
}
