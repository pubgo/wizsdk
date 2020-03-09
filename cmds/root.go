package cmds

import (
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/wizsdk/version"
	"github.com/pubgo/xcmd/xcmd/xcmd_wv"
)

const Service = "WIZ"

// Execute exec
var Execute = xcmd.Init(func(cmd *xcmd.Command) {
	cmd.Version = version.Version

	cmd.AddCommand(
		xcmd_wv.Init(),
	)
})
