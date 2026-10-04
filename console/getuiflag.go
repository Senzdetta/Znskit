// https://github.com/Senzdetta/Znskit

package console

import (
    "os"
    "github.com/Senzdetta/Znskit/module/getuiflag"
    "github.com/Senzdetta/Znskit/utils/invinput"
)

type Getuiflag struct{}
func (c Getuiflag) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    toolName := args[2]
    getuiflag.Fetch(toolName)
}

// Copyright (c) 2026 Senzdetta