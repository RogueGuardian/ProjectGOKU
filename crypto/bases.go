package crypto

import (
    "encoding/base64"
    "fmt"
    "../helpers"
)

func decBase64(inputList []string) {
    var monoOps = []string{}
    var flagOps = []string{"msg"}
    flagVals, flagErr := helpers.ParseInput(monoOps, flagOps, inputList[1:])
    if !flagErr {
        return
    }
    if !helpers.IsFlagIn("msg", flagVals) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Missing msg argument flag |")) 
    }
    decoded, err := base64.StdEncoding.DecodeString(helpers.GetFlagValue("msg", flagVals))
    if err != nil {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Invalid base64 string |")) 
        return
    }
    fmt.Printf("%s %s\n", helpers.PrintColor("bold", "green", "none", "[Base64]"), helpers.PrintColor("bold", "blue", "none", string(decoded)))
}
