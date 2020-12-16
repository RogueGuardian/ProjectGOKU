package helpers

import (
    //"strings"
    "fmt"
)

type FlagPair struct {
    flag string
    value string
}

func ParseInput(monoOps []string, flagOps []string, stringList []string) ([]FlagPair, bool) {
    var flagList []FlagPair
    newBit := []byte("-")[0]
    monoParity := false

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
            fmt.Printf("%s\n", PrintColor("bold", "red", "none", "| ERROR ++ Invalid Mono-Op Flag |"))
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
        //fmt.Printf("%v\n", v)
        if v[0] == newBit && v[1] != newBit {
            fmt.Printf("%s\n", PrintColor("bold", "red", "none", "| ERROR ++ Invalid Argument Flag Prefix |"))
            return flagList, false
        }
        if len(v) == 1 {
            tempVar.value += v+" "
            continue
        }
        if v[0:2] == "--" {
            if k != 0 {
                tempVar.value = tempVar.value[:len(tempVar.value)-1]
                flagList = append(flagList, tempVar)
            }
            if IsStrIn(v[2:len(v)], flagOps) {
                tempVar = FlagPair{
                    flag: v[2:len(v)],
                    value: "",
                }
                continue
            }
            fmt.Printf("%s\n", PrintColor("bold", "red", "none", "| ERROR ++ Invalid Argument Flag |"))
            return flagList, false
        }
        tempVar.value += v+" "
    }
    tempVar.value = tempVar.value[:len(tempVar.value)-1]
    flagList = append(flagList, tempVar)
    return flagList, true
}
