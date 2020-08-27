package helpers

import (
    "strings"
    "fmt"
)

type FlagPair struct {
    flag string
    value string
}

func ParseInput(monoOps []string, flagOps []string, usrString string) ([]FlagPair, bool) {
    var flagList []FlagPair
    newBit := []byte("-")[0]
    monoParity := false
    stringList := strings.Split(usrString, " ")

    // This just removes the first two items <pack> and <fn>
    stringList = stringList[2:]

    // If there are no flags, then just return the empty list
    if len(stringList) == 0 {
        return flagList, true
    }

    // First item should always be all Mono-Ops (No argument flags) [Start with -]
    // They should be in a predefined list, throw error and stop, but not end when unknown flag is present
    // Then comes all flags with paramters (Argument Flags), always space separated [Start with --]
    // Assume flag values go until the next detected flag
    // If not in the predefined list, throw error and stop, but not end when unknown flag is present
    if stringList[0][0] == newBit && stringList[0][1] != newBit {
        for _,v := range stringList[0][1:] {
            v := string(v)
            if IsStrIn(v, monoOps) {
                flagList = append(flagList, FlagPair{
                    flag: v,
                    value: "",
                })
                monoParity = true
                continue
            }
            fmt.Printf("Error ++ Invalid Mono-Op Flag\n")
            return flagList, false
        }
    }
    if len(stringList) == 1 {
        return flagList, true
    }
    
    var tempVar FlagPair
    if monoParity {
        stringList = stringList[1:]
    }
    for k,v := range stringList {
        v := string(v)
        fmt.Printf("%v\n", v)
        if v[0] == newBit && v[1] != newBit {
            fmt.Printf("Error ++ Invalid Argument Flag Prefix\n")
            return flagList, false
        }
        if len(v) == 1 {
            tempVar.value += " "+v
            continue
        }
        if v[0:2] == "--" {
            if k != 0 {
                flagList = append(flagList, tempVar)
            }
            if IsStrIn(v[2:len(v)], flagOps) {
                tempVar = FlagPair{
                    flag: v[1:len(v)],
                    value: "",
                }
                continue
            }
            fmt.Printf("Error ++ Invalid Argument Flag\n")
            return flagList, false
        }
        tempVar.value += " "+v 
    }
    flagList = append(flagList, tempVar)
    return flagList, true
}
