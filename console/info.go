// https://github.com/Senzdetta/Znskit

package console

import (
    "os"
    "github.com/Senzdetta/Znskit/module/info"
    "github.com/Senzdetta/Znskit/utils/invinput"
)

type Info struct{}
func (c Info) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    toolName := args[2]
    info.Show(toolName)
}

// Copyright (c) 2026 Senzdetta