package crypto

import (
    "fmt"
    "strconv"
    "../helpers"
)

func rotCipher(inputList []string) {
    var monoOps = []string{}
    var flagOps = []string{"num", "msg"}

    flagList, err := helpers.ParseInput(monoOps, flagOps, inputList[1:])
    if !err {
        return
    }
    if !helpers.IsFlagIn("msg", flagList) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "ERROR ++ Missing --msg and value"))
        return
    }
    byteList := helpers.GetFlagValue("msg", flagList)
    if helpers.IsFlagIn("num", flagList) {
        intVal, intErr := strconv.Atoi(helpers.GetFlagValue("num", flagList))
        if intErr != nil {
            fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "ERROR ++ Invalid --num parameter"))
            return
        }
        baseRot(byteList, intVal)
        return
    }
    for i := 0; i < 26; i++ {
        baseRot(byteList, i)
    }
}

func baseRot(inputStr string, rotVal int) {
    var newStr string
    var intVal byte
    if rotVal > 0 {
        for bit := range inputStr {
            intVal = inputStr[bit] - byte(rotVal)
            if inputStr[bit] >= 'A' && inputStr[bit] <= 'Z' {
                if intVal < 'A' {
                    intVal += 26
                }
                newStr += string(intVal)
                continue
            } else if inputStr[bit] >= 'a' && inputStr[bit] <= 'z' { 
                if intVal < 'a' {
                    intVal += 26
                }
                newStr += string(intVal)
                continue
            }
            newStr += string(inputStr[bit])
        }
        rotString := helpers.PrintColor("bold", "green", "none", fmt.Sprintf("ROT[%d]", rotVal))
        fmt.Printf("%s %s\n", rotString, helpers.PrintColor("bold", "blue", "none", fmt.Sprintf("%s --> %s", inputStr, newStr)))
        return
    }
}

func atbashCipher(inputList []string) {
    var monoOps = []string{}
    var flagOps = []string{"msg"}
    flagList, err := helpers.ParseInput(monoOps, flagOps, inputList[1:])
    if !err {
        return
    }
    if !helpers.IsFlagIn("msg", flagList) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "ERROR ++ Missing --msg and value"))
        return
    }
    usrStr := helpers.GetFlagValue("msg", flagList) 
    var newStr string
    for i := range usrStr {
        bit := usrStr[i]
        if bit >= 'A' && bit <= 'Z' {
           newStr += string('A'+'Z'-bit) 
           continue
        }
        if bit >= 'a' && bit <= 'z' {
           newStr += string('a'+'z'-bit) 
           continue
        }
        newStr += string(bit)
    }
    fmt.Printf("%s %s\n", helpers.PrintColor("bold", "green", "none", "[ATBASH]"), helpers.PrintColor("bold", "blue", "none", newStr))
}
