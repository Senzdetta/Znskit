// https://github.com/Senzdetta/Znskit

package console

import (
    "github.com/Senzdetta/Znskit/module/help"
)

type Helper struct{}
func (c Helper) Execute(args []string) {
    help.Show()
}

// Copyright (c) 2026 Senzdetta