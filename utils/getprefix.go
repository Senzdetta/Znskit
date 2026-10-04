// https://github.com/Senzdetta/Znskit

package utils

import (
    "os"
)

func GetPrefix() string {
    prefix := os.Getenv("PREFIX")
    if prefix == "" {
        prefix = "/usr"
    }
    return prefix
}

// Copyright (c) 2026 Senzdetta