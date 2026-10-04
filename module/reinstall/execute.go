// https://github.com/Senzdetta/Znskit

package reinstall

import (
    "fmt"
    "github.com/Senzdetta/Znskit/module/install"
    "github.com/Senzdetta/Znskit/module/uninstall"
    "github.com/Senzdetta/Znskit/utils/color"
)

func Execute(toolName string, iFlags []string, uFlags []string) {
    fmt.Printf(
        "%s[*] %sReinstalling: %s%s%s\n",
        color.B, color.N, color.GG, toolName, color.N,
    )

    uninstall.Execute(toolName, uFlags)
    install.Clone(toolName, iFlags)
}

// Copyright (c) 2026 Senzdetta