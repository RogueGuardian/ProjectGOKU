package osint

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "../helpers"
)

func OsintInit() {
    var usrString string
    var stringList []string
    reader := bufio.NewReader(os.Stdin)
    helpers.ClearTerm()
    osintHeader()
    usrString, _ = reader.ReadString('\n')
    usrString = usrString[:len(usrString)-1]
    for strings.ToLower(usrString) != "quit" {
        stringList = strings.Split(usrString, " ")
        switch stringList[0] {
        case "clear":
            helpers.ClearTerm()
            break

		case "userscan":
			usernameScan(stringList[1:])
			break

		case "location":
			location(stringList[1:])
			break

        default:
            fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "ERROR ++ Invalid OSINT function"))

        }
        osintHeader()
        usrString, _ = reader.ReadString('\n')
        usrString = usrString[:len(usrString)-1]
    }
    helpers.ClearTerm()
}

func osintHeader() {
    leftBr := helpers.PrintColor("bold", "blue", "none", "[ ")
    leftBr += helpers.PrintColor("bold", "red", "none", "GOKU")
    leftBr += helpers.PrintColor("bold", "blue", "none", " - ")
    header := helpers.PrintColor("bold", "yellow", "none", "OSINT")
    rightBr := helpers.PrintColor("bold", "blue", "none", " | ")
    fmt.Printf("%s%s%s", leftBr, header, rightBr)
}
