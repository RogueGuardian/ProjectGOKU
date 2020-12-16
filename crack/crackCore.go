package crack

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "../helpers"
)

func CrackInit() {
    var usrString string
    var stringList []string
    reader := bufio.NewReader(os.Stdin)
    helpers.ClearTerm()
    crackHeader()
    usrString, _ = reader.ReadString('\n')
    usrString = usrString[:len(usrString)-1]
    for strings.ToLower(usrString) != "quit" {
        stringList = strings.Split(usrString, " ")
        switch stringList[0] {
        case "clear":
            helpers.ClearTerm()
            break

        case "md5":
            crackMD5(stringList[1:])
            break

        case "sha1":
            crackSHA1(stringList[1:])
            break

        default:
            fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "ERROR ++ Invalid cracking function"))

        }
        crackHeader()
        usrString, _ = reader.ReadString('\n')
        usrString = usrString[:len(usrString)-1]
    }
    helpers.ClearTerm()
}

func crackHeader() {
    leftBr := helpers.PrintColor("bold", "blue", "none", "[ ")
    leftBr += helpers.PrintColor("bold", "red", "none", "GOKU")
    leftBr += helpers.PrintColor("bold", "blue", "none", " - ")
    header := helpers.PrintColor("bold", "yellow", "none", "CRACK")
    rightBr := helpers.PrintColor("bold", "blue", "none", " | ")
    fmt.Printf("%s%s%s", leftBr, header, rightBr)
}
