package cmds

import (
	"github.com/pubgo/wizsdk/version"
	"github.com/pubgo/xcmd/xcmd"
	"github.com/pubgo/xcmd/xcmd/xcmd_wv"
)

const Service = "wiz"

// Execute exec
var Execute = xcmd.Init(func(cmd *xcmd.Command) {
	cmd.Version = version.Version
	cmd.Use = Service

	cmd.AddCommand(
		xcmd_wv.Init(),
	)
})
