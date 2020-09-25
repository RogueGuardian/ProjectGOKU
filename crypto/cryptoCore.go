package crypto

import (
    "fmt"
    "../helpers"
)

func CryptoStart(inputItems []string) {
    var fn string
    fn = inputItems[0]
    switch fn {
    case "rot":
        rotCipher(inputItems)
        break
    case "atbash":
        atbashCipher(inputItems)
        break
    default:
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "ERROR ++ Invalid crypto function"))
    }
}
