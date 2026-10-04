// https://github.com/Senzdetta/Znskit

package console

import (
    "os"
    "strings"
    "github.com/Senzdetta/Znskit/module/uninstall"
    "github.com/Senzdetta/Znskit/utils/invinput"
)

type Uninstall struct{}
func (c Uninstall) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    toolName := args[2]
    var uFlags []string

    for i := 3; i < len(args); i++ {
        if args[i] == "--uflag" && i+1 < len(args) {
            rawFlags := args[i+1]
            if strings.TrimSpace(rawFlags) != "" {
                uFlags = strings.Fields(rawFlags)
            }
            break
        }
    }

    uninstall.Execute(toolName, uFlags)
}

// Copyright (c) 2026 Senzdetta