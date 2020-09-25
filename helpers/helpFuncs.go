package helpers

import (
    "fmt"
    "os"
)

func ClearMemory() {
    os.RemoveAll("./GokuTempDir")
}

func ClearTerm() {
    fmt.Printf("\033[2J")
    fmt.Printf("\033[H")
}

func IsStrIn(item string, list []string) bool {
    for i := range list {
        if list[i] == item {
            return true
        }
    }
    return false
}

func IsIntIn(item int, list []int) bool {
    for i := range list {
        if list[i] == item {
            return true
        }
    }
    return false
}

func IsFlagIn(flag string, flagList []FlagPair) bool {
    for i := range flagList {
        if flagList[i].flag == flag {
            return true
        }
    }
    return false
}

func GetFlagValue(flag string, flagList []FlagPair) string {
    for i := range flagList {
        if flagList[i].flag == flag {
            return flagList[i].value
        }
    }
    return "ERROR"
}
