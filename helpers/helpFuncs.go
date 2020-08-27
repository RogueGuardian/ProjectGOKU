package helpers

import (
)

func IsStrIn(item string, list []string) bool {
    for i := range list {
        if list[i] == item {
            return true
        }
    }
    return false
}
