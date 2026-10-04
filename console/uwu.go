// https://github.com/Senzdetta/Znskit

package console

import (
    "time"
    "fmt"
    "github.com/Senzdetta/Znskit/module/uwu"
)

type UWU struct{}
func (c UWU) Execute(args []string) {
    fmt.Printf("\x1b[?25l")
    uwu.Show(5 * time.Second)
    fmt.Printf("\x1b[?25h")

    fmt.Println()
}

// Copyright (c) 2026 Senzdetta