package gambler

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "../helpers"
)

type buffer struct {
    stringList [][]string
    delimeter string
}

// Callable Functions
func (B *buffer) count(stringList []string){
    var total uint64
    var rest string
    col, iErr := strconv.Atoi(stringList[0])
    if iErr != nil {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid integer specified |"))
        return
    }
    cond := stringList[1]
    extra := stringList[2:]
    if len(extra) == 1 {
       rest = extra[0] 
    } else if strings.Contains(extra[0], "\"") {
        if strings.Contains(extra[len(extra)-1], "\"") {
            rest = strings.Join(extra, " ")
            rest = rest[1:len(rest)-1]
        }
    }
    if col >= len(B.stringList[0]) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Index out of bounds |"))
        return
    }
    for i := range B.stringList {
        switch cond {
        case "==":
            if B.stringList[i][col] == rest {         
                total += 1
            }
            break
        case "!=":
            if B.stringList[i][col] != rest {         
                total += 1
            }
            break
        case ">=":
            if B.stringList[i][col] >= rest {         
                total += 1
            }
            break
        case "<=":
            if B.stringList[i][col] <= rest {         
                total += 1
            }
            break
        case "<":
            if B.stringList[i][col] < rest {         
                total += 1
            }
            break
        case ">":
            if B.stringList[i][col] > rest {
                total += 1
            }
            break
        case "missing":
            if !strings.Contains(B.stringList[i][col], rest) { 
                total += 1
            }
            break
        case "contains":
            if strings.Contains(B.stringList[i][col], rest) { 
                total += 1
            }
            break
        default:
            fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid comparison |"))
            return
        }
    }
    fmt.Printf("%s\n", helpers.PrintColor("bold", "blue", "none", fmt.Sprintf("%d", total)))
}


func (B *buffer) dump(){
    B.stringList = nil 
}

func (B *buffer) load(filename []string) {
    fileAAH := strings.Join(filename, " ")
    file, err := os.Open(fileAAH)
    defer file.Close()
    if err != nil {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid file |"))
        return
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        newArr := append([]string{}, scanner.Text()[:len(scanner.Text())])
        B.stringList = append(B.stringList, newArr)
    }
    fmt.Printf("%s\n", helpers.PrintColor("bold", "green", "none", "| File successfully loaded |"))
}

func (B *buffer) merge(stringList []string) {
    start, sErr := strconv.Atoi(stringList[0])
    end, eErr := strconv.Atoi(stringList[1])
    if sErr != nil || eErr != nil {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "|ERROR ++ Invalid integers specified|"))
        return
    }
    if start < 0 || start >= len(B.stringList[0]) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid start position |"))
        return
    }
    if end >= len(B.stringList[0]) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid end position |"))
        return
    }
    var tempList []string
    var tempString string
    for i := range B.stringList {
        tempList = nil
        tempString = ""
        for j := range B.stringList[i] {
            if j == end+1 {
                tempList = append(tempList, tempString[:len(tempString)-1])
                tempList = append(tempList, B.stringList[i][j])
            } else if j == len(B.stringList[i])-1 && j == end {
                tempString += B.stringList[i][j]
                tempList = append(tempList, tempString)
            } else if j >= start && j <= end {
                tempString += B.stringList[i][j]+B.delimeter
            } else {
                tempList = append(tempList, B.stringList[i][j])
            }
        }
        B.stringList[i] = tempList
    }
}

func (B *buffer) preview() {
    var leftBuff int
    var rightBuff int
    if len(B.stringList) == 0 {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "yellow", "none", "| Buffer is currently empty |"))
        return
    }
    for i := range B.stringList[0] {
        leftBuff = len(B.stringList[0][i])/2
        if len(B.stringList[0][i]) % 2 == 0 && leftBuff != 0 {
            leftBuff -= 1
        }
        rightBuff = len(B.stringList[0][i]) - (leftBuff + (i/10) + 1)
        if rightBuff < 0 {
            rightBuff = 0
        }
        fmt.Printf("[%s%d%s]", strings.Repeat(" ", leftBuff), i, strings.Repeat(" ", rightBuff))
    }
    fmt.Printf("\n")
    num := 10
    if len(B.stringList) < 10 {
        num = len(B.stringList)
    }
    for i := 0; i < num; i++ {
        for j := range B.stringList[i] {
            fmt.Printf("[%s]", B.stringList[i][j])
        }
        fmt.Printf("\n")
    }
}

func (B *buffer) buffPrint(inputList []string) {
    // This creates a new buffer and then removes it
    // Maybe in future create new helper functions for print not append
    newB := buffer{}
    newB.stringList = B.stringList
    newB.delimeter = B.delimeter
    newB.remove(inputList)

    // This SHOULD stay the same, but might change
    var leftBuff int
    var rightBuff int
    if len(B.stringList) == 0 {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "yellow", "none", "| Buffer is currently empty |"))
        return
    }
    for i := range B.stringList[0] {
        leftBuff = len(B.stringList[0][i])/2
        if len(B.stringList[0][i]) % 2 == 0 && leftBuff != 0 {
            leftBuff -= 1
        }
        rightBuff = len(B.stringList[0][i]) - (leftBuff + (i/10) + 1)
        if rightBuff < 0 {
            rightBuff = 0
        }
        fmt.Printf("[%s%d%s]", strings.Repeat(" ", leftBuff), i, strings.Repeat(" ", rightBuff))
    }
    fmt.Printf("\n")
    num := len(B.stringList)
    for i := 0; i < num; i++ {
        for j := range B.stringList[i] {
            fmt.Printf("[%s]", B.stringList[i][j])
        }
        fmt.Printf("\n")
    }
    newB.dump()
}

func (B *buffer) remove(inputList []string) {
    remType := inputList[0]
    switch remType {
    case "c":
        idx, iErr := strconv.Atoi(inputList[1])
        if iErr != nil {
            fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid integer specified |"))
            break
        }
        B.removeColumn(idx)
    case "r":
        idx, iErr := strconv.Atoi(inputList[1])
        if iErr != nil {
            fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid integer specified |"))
            break
        }
        B.removeRows(idx, inputList[2], inputList[3:])
        break 
    default:
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid remove character. Must be \"c\" or \"r\" |"))
        break
    }
}

func (B *buffer) save(filename string) {
    file, err := os.Create(filename)
    if err != nil {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid filename |"))
        return
    }
    defer file.Close()
    for i := range B.stringList {
        for j := range B.stringList[i] {
            if j != len(B.stringList[i])-1 {
                fmt.Fprintf(file, "%s%s", B.stringList[i][j], B.delimeter)
                continue
            }
            fmt.Fprintf(file, "%s\n", B.stringList[i][j])
        }
    }
    fmt.Printf("%s\n", helpers.PrintColor("bold", "green", "none", "| File successfully saved |"))
    return
}

func (B *buffer) sort(inputItems []string) {
    intV, intE := strconv.Atoi(inputItems[0])
    if intE != nil {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid integer specified |"))
        return
    }
    if intV >= len(B.stringList[0]) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid column specified |"))
        return
    }
    if !helpers.IsStrIn(inputItems[1], []string{"l", "s", "n", "m"}) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid ordering specified |"))
        return
    }
    for i := range B.stringList {
        for j := i; j < len(B.stringList); j++ {
            switch inputItems[1] {
            case "s":
                if B.stringList[i][intV] > B.stringList[j][intV] {
                    tempVar := B.stringList[i]
                    B.stringList[i] = B.stringList[j]
                    B.stringList[j] = tempVar
                }
            case "l":
                if B.stringList[i][intV] < B.stringList[j][intV] {
                    tempVar := B.stringList[i]
                    B.stringList[i] = B.stringList[j]
                    B.stringList[j] = tempVar
                }

            case "n":
                fNum, fErr := strconv.ParseFloat(B.stringList[i][intV], 64)
                sNum, sErr := strconv.ParseFloat(B.stringList[j][intV], 64)
                if fErr != nil || sErr != nil {
                    fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid numerical field |")) 
                    continue
                }
                if fNum < sNum {
                    tempVar := B.stringList[i]
                    B.stringList[i] = B.stringList[j]
                    B.stringList[j] = tempVar
                }

            case "m":
                fNum, fErr := strconv.ParseFloat(B.stringList[i][intV], 64)
                sNum, sErr := strconv.ParseFloat(B.stringList[j][intV], 64)
                if fErr != nil || sErr != nil {
                    fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid numerical field |")) 
                    continue
                }
                if fNum > sNum {
                    tempVar := B.stringList[i]
                    B.stringList[i] = B.stringList[j]
                    B.stringList[j] = tempVar
                }

            }
        }
    }
}

func (B *buffer) splitLines(del string) {
    if B.delimeter == "" {
        B.delimeter = del 
    }
    for i := range B.stringList {
        B.stringList[i] = strings.Split(B.stringList[i][0], del)
    }
}

// Helper Functions
func (B *buffer) removeColumn(col int) {
    if col >= len(B.stringList[0]) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Index out of bounds |"))
    }
    for i := range B.stringList {
        B.stringList[i] = append(B.stringList[i][:col], B.stringList[i][col+1:]...) 
    }
}

func (B *buffer) removeRows(col int, cond string, extra []string) {
    var newBuffer [][]string
    var rest string
    if len(extra) == 1 {
       rest = extra[0] 
    } else if strings.Contains(extra[0], "\"") {
        if strings.Contains(extra[len(extra)-1], "\"") {
            rest = strings.Join(extra, " ")
            rest = rest[1:len(rest)-1]
        }
    }
    if col >= len(B.stringList[0]) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Index out of bounds |"))
        return
    }
    //fmt.Printf("%v %v %v\n", col, cond, rest)
    for i := range B.stringList {
        switch cond {
        case "==":
            if B.stringList[i][col] != rest {         
                newBuffer = append(newBuffer, B.stringList[i])
            }
            break
        case "!=":
            if B.stringList[i][col] == rest {         
                newBuffer = append(newBuffer, B.stringList[i])
            }
            break
        case ">=":
            if B.stringList[i][col] < rest {         
                newBuffer = append(newBuffer, B.stringList[i])
            }
            break
        case "<=":
            if B.stringList[i][col] > rest {         
                newBuffer = append(newBuffer, B.stringList[i])
            }
            break
        case "<":
            if B.stringList[i][col] >= rest {         
                newBuffer = append(newBuffer, B.stringList[i])
            }
            break
        case ">":
            if B.stringList[i][col] <= rest {
                newBuffer = append(newBuffer, B.stringList[i])
            }
            break
        case "contains":
            if !strings.Contains(B.stringList[i][col], rest) { 
                newBuffer = append(newBuffer, B.stringList[i])
            }
            break
        case "missing":
            if strings.Contains(B.stringList[i][col], rest) { 
                newBuffer = append(newBuffer, B.stringList[i])
            }
            break
        default:
            fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid comparison |"))
            return
        }
    }
    B.stringList = newBuffer
}
