package helpers

import (
    "strconv"
    "strings"
    "fmt"
)

const (
    BLACK = iota
    RED
    GREEN
    YELLOW
    BLUE
    PURPLE
    CYAN
    WHITE
)

const (
    REGULAR = iota
    BOLD
    NA0
    ITALIC
    UNDERLINE
    FLASHING
    NA1
    INVERT
    NA2
    STRIKE
)

func PrintColor(style string, fg string, bg string, usrString string) string {
    var styleInt int
    var fgInt int
    var bgInt int

    switch strings.ToLower(style) {
    case "regular":
        styleInt = REGULAR
        break
    case "bold":
        styleInt = BOLD
        break
    case "italic":
        styleInt = ITALIC
        break
    case "underline":
        styleInt = UNDERLINE
        break
    case "flashing":
        styleInt = FLASHING
        break
    case "inverted":
        styleInt = INVERT
        break
    case "strikethrough":
        styleInt = STRIKE
        break
    default:
        styleInt = -1
        break
    }

    switch strings.ToLower(fg) {
    case "black":
        fgInt = BLACK
        break
    case "red":
        fgInt = RED
        break
    case "green":
        fgInt = GREEN
        break
    case "yellow":
        fgInt = YELLOW
        break
    case "blue":
        fgInt = BLUE
        break
    case "purple":
        fgInt = PURPLE
        break
    case "cyan":
        fgInt = CYAN
        break
    case "white":
        fgInt = WHITE
        break
    default:
        fgInt = -1
        break
    }


    switch strings.ToLower(bg) {
    case "black":
        bgInt = BLACK
        break
    case "red":
        bgInt = RED
        break
    case "green":
        bgInt = GREEN
        break
    case "yellow":
        bgInt = YELLOW
        break
    case "blue":
        bgInt = BLUE
        break
    case "purple":
        bgInt = PURPLE
        break
    case "cyan":
        bgInt = CYAN
        break
    case "white":
        bgInt = WHITE
        break
    default:
        bgInt = -1
        break
    }

    return fmt.Sprintf("\033[%sm%s\033[0m", genColorCode(styleInt, fgInt, bgInt), usrString)
}

func genColorCode(style int, fg int, bg int) string {
    var setStr string

    setStr += strconv.Itoa(style)
    if fg != -1 {
        setStr += ";"+strconv.Itoa(fg+30)
    }
    if bg != -1 {
        setStr += ";"+strconv.Itoa(bg+40)
    }

    return setStr
}
