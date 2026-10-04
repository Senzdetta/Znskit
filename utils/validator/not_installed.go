// https://github.com/Senzdetta/Znskit

package validator

func NotInstalled(toolName string) bool {
    return !Installed(toolName)
}

// Copyright (c) 2026 Senzdetta