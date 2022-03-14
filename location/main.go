package main

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

func main() {
    fmt.Println(GetCurrPath())
}

func GetCurrPath() string {
    file, _ := exec.LookPath("./main,go")
    path, _ := filepath.Abs(file)
    index := strings.LastIndex(path, string(os.PathSeparator))
    ret := path[:index]
    return ret
}