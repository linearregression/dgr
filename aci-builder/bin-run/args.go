package main

import (
	"flag"
	"fmt"
	"github.com/appc/spec/schema/types"
	"github.com/blablacar/dgr/bin-dgr/common"
	rktcommon "github.com/coreos/rkt/common"
	"github.com/n0rad/go-erlog/logs"
	"os"
)

var cliDebugFlag bool

func init() {

	var discardString string
	var discardBool bool
	var discardNetlist rktcommon.NetList

	flag.BoolVar(&cliDebugFlag, "debug", false, "Run in debug mode")

	// The following flags need to be supported by stage1 according to
	// https://github.com/coreos/rkt/blob/master/Documentation/devel/stage1-implementors-guide.md
	// TODO: either implement functionality or give not implemented warnings
	flag.Var(&discardNetlist, "net", "Setup networking")
	flag.BoolVar(&discardBool, "interactive", true, "The pod is interactive")
	flag.StringVar(&discardString, "mds-token", "", "MDS auth token")
	flag.StringVar(&discardString, "local-config", rktcommon.DefaultLocalConfigDir, "Local config path")
}

func ProcessArgsAndReturnPodUUID() *types.UUID {
	flag.Parse()

	if cliDebugFlag {
		logs.SetLevel(logs.DEBUG)
	}
	if lvlStr := os.Getenv(common.EnvLogLevel); lvlStr != "" {
		lvl, err := logs.ParseLevel(lvlStr)
		if err != nil {
			fmt.Printf("Unknown log level : %s", lvlStr)
			os.Exit(1)
		}
		logs.SetLevel(lvl)
	}

	arg := flag.Arg(0)
	uuid, err := types.NewUUID(arg)
	if err != nil {
		logs.WithE(err).WithField("content", arg).Fatal("UUID is missing or malformed")
	}
	return uuid
}
