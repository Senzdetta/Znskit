// https://github.com/Senzdetta/Znskit

package console

import (
    "os"
    "strings"
    "github.com/Senzdetta/Znskit/module/search"
    "github.com/Senzdetta/Znskit/utils/invinput"
)

type Search struct{}
func (c Search) Execute(args []string) {
    if len(args) < 3 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    keyword := strings.Join(args[2:], " ")
    search.Searcher(keyword)
}

// Copyright (c) 2026 Senzdetta