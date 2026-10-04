// https://github.com/Senzdetta/Znskit

package version

import (
    "fmt"
    "github.com/Senzdetta/Znskit/utils/color"
)

const (
    name = "Znskit"
    version = "v0.1.04102026"
    creator = "Senzdetta"
    homepage = "https://github.com/Senzdetta/Znskit"
)

func Show() {
    fmt.Printf(
        "%sName: %s%s%s\n",
        color.N, color.GG, name, color.N,
    )

    fmt.Printf(
        "%sVersion: %s%s%s\n",
        color.N, color.GG, version, color.N,
    )

    fmt.Printf(
        "%sCreator: %s%s%s\n",
        color.N, color.GG, creator, color.N,
    )

    fmt.Printf(
        "%sHomepage: %s%s%s\n",
        color.N, color.GG, homepage, color.N,
    )
}

// Copyright (c) 2026 Senzdetta