// https://github.com/Senzdetta/Znskit

package console

import (
    "github.com/Senzdetta/Znskit/module/version"
)

type Version struct{}
func (c Version) Execute(args []string) {
    version.Show()
}

// Copyright (c) 2026 Senzdetta