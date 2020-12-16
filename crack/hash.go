package crack

import (
    "bufio"
    "crypto/md5"
    "crypto/sha1"
    "fmt"
    "os"
    "sync"
    "../helpers"
)

var monoOps = []string{}
var flagOps = []string{"file", "hash", "wordlist"}
// filename, hash, wordlist

func crackMD5(inputList []string) {
    flagList, err := helpers.ParseInput(monoOps, flagOps, inputList[1:])
    if !err {
        return
    }
    if !helpers.IsFlagIn("hash", flagList) && !helpers.IsFlagIn("file", flagList) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ No hash value or file specified |"))
        return
    }
    if !helpers.IsFlagIn("wordlist", flagList) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ No wordlist specified |"))
        return
    }
    if helpers.IsFlagIn("hash", flagList) {
        hashVal := helpers.GetFlagValue("hash", flagList)
        wordVal := helpers.GetFlagValue("wordlist", flagList)
        fmt.Printf("%s\n", md5Hash(hashVal, wordVal))
        return
    }
    // Run Wordlist
    var wg sync.WaitGroup
    wordVal := helpers.GetFlagValue("wordlist", flagList)
    hashFile, hashErr := os.Open(helpers.GetFlagValue("file", flagList))
    if hashErr != nil {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Hash location invalid |"))
        return
    }
    defer hashFile.Close()
    hashScan := bufio.NewScanner(hashFile)
    for hashScan.Scan() {
        wg.Add(1)
        go func(text string){
            fmt.Printf("%s\n", md5Hash(text, wordVal))
            wg.Done()
        }(hashScan.Text())
    }
    wg.Wait()
    return
}

func crackSHA1(inputList []string) {
    flagList, err := helpers.ParseInput(monoOps, flagOps, inputList[1:])
    if !err {
        return
    }
    if !helpers.IsFlagIn("hash", flagList) && !helpers.IsFlagIn("file", flagList) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ No hash location specified |"))
        return
    }
    if !helpers.IsFlagIn("wordlist", flagList) {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ No wordlist specified |"))
        return
    }
    if helpers.IsFlagIn("hash", flagList) {
        hashVal := helpers.GetFlagValue("hash", flagList)
        wordVal := helpers.GetFlagValue("wordlist", flagList)
        fmt.Printf("%s\n", md5Hash(hashVal, wordVal))
        return
    }
    // Run Wordlist
    var wg sync.WaitGroup
    wordVal := helpers.GetFlagValue("wordlist", flagList)
    hashFile, hashErr := os.Open(helpers.GetFlagValue("file", flagList))
    if hashErr != nil {
        fmt.Printf("%s\n", helpers.PrintColor("bold", "red", "none", "| ERROR ++ Hash location invalid |"))
        return
    }
    defer hashFile.Close()
    hashScan := bufio.NewScanner(hashFile)
    for hashScan.Scan() {
        wg.Add(1)
        go func(text string){
            fmt.Printf("%s\n", sha1Hash(text, wordVal))
            wg.Done()
        }(hashScan.Text())
    }
    wg.Wait()
    return
}

func sha1Hash(hash string, wordlist string) string {
    file, err := os.Open(wordlist)
    if err != nil {
        return helpers.PrintColor("bold", "red", "none", "| ERROR ++ Wordlist not valid |")
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        hashArr := sha1.Sum([]byte(scanner.Text()))
        newWord := fmt.Sprintf("%x", hashArr)
        if newWord == hash {
            return helpers.PrintColor("bold", "blue", "none", fmt.Sprintf("%s --> %s", hash, scanner.Text()))
        }
    }
    return helpers.PrintColor("italic", "red", "none", fmt.Sprintf("%s --> HASH NOT FOUND", hash))
}

func md5Hash(hash string, wordlist string) string {
    file, err := os.Open(wordlist)
    if err != nil {
        return helpers.PrintColor("bold", "red", "none", "| ERROR ++ Wordlist not valid |")
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        hashArr := md5.Sum([]byte(scanner.Text()))
        newWord := fmt.Sprintf("%x", hashArr)
        if newWord == hash {
            return helpers.PrintColor("bold", "blue", "none", fmt.Sprintf("%s --> %s", hash, scanner.Text()))
        }
    }
    return helpers.PrintColor("italic", "red", "none", fmt.Sprintf("%s --> HASH NOT FOUND", hash))
}
