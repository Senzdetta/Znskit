// https://github.com/Senzdetta/Znskit

package console

import (
    "github.com/Senzdetta/Znskit/module/listins"
)

type ListInstalled struct{}
func (c ListInstalled) Execute(args []string) {
    listins.Show()
}

// Copyright (c) 2026 Senzdetta