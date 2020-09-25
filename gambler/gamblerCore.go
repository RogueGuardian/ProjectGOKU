package gambler

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "../helpers"
)

func GamblerInit() {
    var usrString string
    var stringList []string
    var Buff buffer
    reader := bufio.NewReader(os.Stdin)
    helpers.ClearTerm()
    gamblerHeader()
    usrString, _ = reader.ReadString('\n')
    usrString = usrString[:len(usrString)-1]
    for strings.ToLower(usrString) != "quit" {
        stringList = strings.Split(usrString, " ")
        switch stringList[0] {
        case "clear":
            helpers.ClearTerm()
            break

        case "count":
            if len(stringList) < 4 {
                fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid input parameters|"))
                break
            }
            Buff.count(stringList[1:])
            break

        case "dump":
            Buff.dump()
            break

        case "load":
            if len(stringList) < 2 {
                fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid input parameters|"))
                break
            }
            Buff.load(stringList[1:])
            break

        case "merge":
            if len(stringList) < 3 {
                fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid input parameters|"))
                break
            }
            Buff.merge(stringList[1:2])
            break
            
        case "preview":
            Buff.preview() 
            break

        case "print":
            if len(stringList) < 2 {
                fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid input parameters|"))
                break 
            }
            Buff.buffPrint(stringList[1:])
            break
            
        case "remove":
            if len(stringList) < 2 {
                fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid input parameters|"))
                break
            }
            Buff.remove(stringList[1:])
            break
            
        case "save":
            if len(stringList) < 2 {
                fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid input parameters|"))
                break
            }
            Buff.save(stringList[1])

        case "sort":
            if len(stringList) < 3 {
                fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid input parameters|"))
                break
            }
            Buff.sort(stringList[1:])
            break
            
        case "split":
            if len(stringList) < 2 {
                fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid input parameters|"))
                break
            }
            if stringList[1] == "" {
                Buff.splitLines(" ")
            } else {
                Buff.splitLines(stringList[1])
            }
            break
            
        default:
            errorVar := helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid command |")
            fmt.Printf("%s", errorVar)
            reader.ReadString('\n')
            break
        }
        gamblerHeader()
        usrString, _ = reader.ReadString('\n')
        usrString = usrString[:len(usrString)-1]
    }
    helpers.ClearTerm()
}

func gamblerHeader() {
    leftBr := helpers.PrintColor("bold", "blue", "none", "[ ")
    header := helpers.PrintColor("bold", "yellow", "none", "GAMBLER")
    rightBr := helpers.PrintColor("bold", "blue", "none", " | ")
    fmt.Printf("%s%s%s", leftBr, header, rightBr)
}
