package crack

import (
    "../helpers"
    "fmt"
)

func CrackStart(inputItems []string) {
    var fn string
    fn = inputItems[0]
    switch fn {
    case "md5":
        crackMD5(inputItems)
        break
    case "sha1":
        crackSHA1(inputItems)
        break
    default:
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "ERROR ++ Invalid cracking function"))
    }
}
