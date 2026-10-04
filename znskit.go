// https://github.com/Senzdetta/Znskit

package main

import (
    "os"
    "strings"
    "github.com/Senzdetta/Znskit/console"
)

func main() {
    args := os.Args[1:]
    input := strings.Join(args, " ")
    console.ZnsConsole(input)
}

// Copyright (c) 2026 Senzdetta