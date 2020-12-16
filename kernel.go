package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "./crack"
    "./crypto"
    "./gambler"
    "./helpers"
)

func main() {
    var usrString string
    var stringList []string
    reader := bufio.NewReader(os.Stdin)
    //tempDirErr := os.Mkdir("./GokuTempDir", 0755)
    //if tempDirErr != nil {
    //    panic(tempDirErr)
    //}
    helpers.ClearTerm()
    header()
    usrString, _ = reader.ReadString('\n')
    usrString = usrString[:len(usrString)-1]
    for strings.ToLower(usrString) != "quit" {
        stringList = strings.Split(usrString, " ")
        if len(stringList) == 0 {
            continue
        }
        switch stringList[0] {
        case "clear":
            helpers.ClearTerm()
            break
        case "crack":
            crack.CrackInit()
            break
        case "crypto":
            crypto.CryptoInit()
            break
        case "exit":
            errorVar := helpers.PrintColor("bold", "purple", "none", "| ERROR ++ Did you mean quit? |")
            fmt.Printf("%s", errorVar)
            reader.ReadString('\n')
            break
        case "gambler":
            gambler.GamblerInit()
        default:
            errorVar := helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid pack |")
            fmt.Printf("%s", errorVar)
            reader.ReadString('\n')
            break
        }
        header()
        usrString, _ = reader.ReadString('\n')
        usrString = usrString[:len(usrString)-1]
    }
    helpers.ClearTerm()
    err := os.RemoveAll("./GokuTempDir")
    if err != nil {
        panic(err)
    }
}

func header() {
    leftBr := helpers.PrintColor("bold", "blue", "none", "[")
    rightBr := helpers.PrintColor("bold", "blue", "none", "|")
    goku := helpers.PrintColor("bold", "red", "none", " GOKU ")
    fmt.Printf("%s%s%s ", leftBr, goku, rightBr)
}
