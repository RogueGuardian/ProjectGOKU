package crypto

import (
    "fmt"
    "strconv"
    "../helpers"
)

/*
Input: inputList = List of strings parsed by spaces
Output: Rotated string in [num] ways
Return: NA
*/
func rotCipher(inputList []string) {
    var monoOps = []string{}
    var flagOps = []string{"num"}

    if len(inputList) == 0 {
        return
    }

    flagList, err := helpers.ParseInput(monoOps, flagOps, inputList[:])
    if !err {
        return
    }
    byteList := inputList[0]
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

/*
Input: inputStr = Input string to be rotated; rotVal = Value to be rotated by
Output: Prints the rotated string in a formated way
Return: NA
*/
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

/*
Input: inputList = List of string[s] to be atbashed
Output: Atbash strings printed to the terminal
Return: NA
*/
func atbashCipher(inputList []string) {

    if len(inputList) == 0 {
        return
    }

    usrStr := inputList[0]
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
