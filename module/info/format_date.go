// https://github.com/Senzdetta/Znskit

package info

import (
    "time"
)

func formatDate(isoString string) string {
    t, err := time.Parse(time.RFC3339, isoString)
    if err != nil {
        return isoString
    }
    return t.Format("2006-01-02")
}

// Copyright (c) 2026 Senzdetta