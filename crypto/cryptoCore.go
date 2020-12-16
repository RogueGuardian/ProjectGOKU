package crypto

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "../helpers"
)

func CryptoInit() {
    var usrString string
    var stringList []string
    reader := bufio.NewReader(os.Stdin)
    helpers.ClearTerm()
    cryptoHeader()
    usrString, _ = reader.ReadString('\n')
    usrString = usrString[:len(usrString)-1]
    for strings.ToLower(usrString) != "quit" {
        stringList = strings.Split(usrString, " ")
        switch stringList[0] {
        case "clear":
            helpers.ClearTerm()
            break

        case "rot":
            rotCipher(stringList[1:])
            break

        case "atbash":
            atbashCipher(stringList[1:])
            break

        default:
            fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "ERROR ++ Invalid crypto function"))

        }
        cryptoHeader()
        usrString, _ = reader.ReadString('\n')
        usrString = usrString[:len(usrString)-1]
    }
    helpers.ClearTerm()
}

func cryptoHeader() {
    leftBr := helpers.PrintColor("bold", "blue", "none", "[ ")
    leftBr += helpers.PrintColor("bold", "red", "none", "GOKU")
    leftBr += helpers.PrintColor("bold", "blue", "none", " - ")
    header := helpers.PrintColor("bold", "yellow", "none", "CRYPTO")
    rightBr := helpers.PrintColor("bold", "blue", "none", " | ")
    fmt.Printf("%s%s%s", leftBr, header, rightBr)
}
