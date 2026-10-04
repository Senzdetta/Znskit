// https://github.com/Senzdetta/Znskit

package console

import (
    "os"
    "strings"
    "github.com/Senzdetta/Znskit/module/reinstall"
    "github.com/Senzdetta/Znskit/utils/invinput"
)

type Reinstall struct{}
func (c Reinstall) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    toolName := args[2]
    var iFlags []string
    var uFlags []string

    for i := 3; i < len(args); i++ {
        if args[i] == "--iflag" && i+1 < len(args) {
            rawFlags := args[i+1]
            if strings.TrimSpace(rawFlags) != "" {
                iFlags = strings.Fields(rawFlags)
            }
            i++
        } else if args[i] == "--uflag" && i+1 < len(args) {
            rawFlags := args[i+1]
            if strings.TrimSpace(rawFlags) != "" {
                uFlags = strings.Fields(rawFlags)
            }
            i++
        }
    }

    reinstall.Execute(toolName, iFlags, uFlags)
}

// Copyright (c) 2026 Senzdetta