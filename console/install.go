// https://github.com/Senzdetta/Znskit

package console

import (
    "os"
    "strings"
    "github.com/Senzdetta/Znskit/module/install"
    "github.com/Senzdetta/Znskit/utils/invinput"
)

type Install struct{}
func (c Install) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    toolName := args[2]
    var iFlags []string

    for i := 3; i < len(args); i++ {
        if args[i] == "--iflag" && i+1 < len(args) {
            rawFlags := args[i+1]
            if strings.TrimSpace(rawFlags) != "" {
                iFlags = strings.Fields(rawFlags)
            }
            break
        }
    }

    install.Clone(toolName, iFlags)
}

// Copyright (c) 2026 Senzdetta