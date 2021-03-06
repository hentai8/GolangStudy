package main

import (
    "bytes"
    "fmt"
    "strings"
)

func main() {
    // 字符串和字符切片的操作基本一致
    // 包含
    fmt.Println(strings.Contains("Golang", "Go"))
    fmt.Println(strings.Contains("Golang", "go"))
    fmt.Println(strings.Contains("Golang", ""))
    fmt.Println(strings.Contains("", ""))

    fmt.Println(bytes.Contains([]byte("Golang"), []byte("Go"))) // true
    fmt.Println(bytes.Contains([]byte("Golang"), []byte("go"))) // false
    fmt.Println(bytes.Contains([]byte("Golang"), []byte("")))   // true
    fmt.Println(bytes.Contains([]byte("Golang"), nil))   // true
    fmt.Println(bytes.Contains([]byte("Golang"), []byte{}))     // true
    fmt.Println(bytes.Contains(nil, nil))             // true


    // 交集
    fmt.Println(strings.ContainsAny("Golang", "java"))   // true
    fmt.Println(strings.ContainsAny("Golang", "python")) // true
    fmt.Println(strings.ContainsAny("Golang", "c"))      // false
    fmt.Println(strings.ContainsAny("Golang", ""))       // false
    fmt.Println(strings.ContainsAny("", ""))             // false

    fmt.Println(bytes.ContainsAny([]byte("Golang"), "java")) // true
    fmt.Println(bytes.ContainsAny([]byte("Golang"), "c"))    // false
    fmt.Println(bytes.ContainsAny([]byte("Golang"), ""))     // false
    fmt.Println(bytes.ContainsAny(nil, ""))                  // false

    //HasPrefix 函数用于判断第二个参数代表的字符串 / 字节切片是否是第一个参数的前缀，同理，HasSuffix 函数则用于判断第二个参数是否是第一个参数的后缀。
    fmt.Println(strings.HasPrefix("Golang", "Go"))     // true
    fmt.Println(strings.HasPrefix("Golang", "Golang")) // true
    fmt.Println(strings.HasPrefix("Golang", "lang"))   // false
    fmt.Println(strings.HasPrefix("Golang", ""))       // true
    fmt.Println(strings.HasPrefix("", ""))             // true

    fmt.Println(strings.HasSuffix("Golang", "Go"))     // false
    fmt.Println(strings.HasSuffix("Golang", "Golang")) // true
    fmt.Println(strings.HasSuffix("Golang", "lang"))   // true
    fmt.Println(strings.HasSuffix("Golang", ""))       // true
    fmt.Println(strings.HasSuffix("", ""))             // true

    fmt.Println(bytes.HasPrefix([]byte("Golang"), []byte("Go")))     // true
    fmt.Println(bytes.HasPrefix([]byte("Golang"), []byte("Golang"))) // true
    fmt.Println(bytes.HasPrefix([]byte("Golang"), []byte("lang")))   // false
    fmt.Println(bytes.HasPrefix([]byte("Golang"), []byte{}))         // true
    fmt.Println(bytes.HasPrefix([]byte("Golang"), nil))              // true
    fmt.Println(bytes.HasPrefix(nil, nil))                           // true

    fmt.Println(bytes.HasSuffix([]byte("Golang"), []byte("Go")))     // false
    fmt.Println(bytes.HasSuffix([]byte("Golang"), []byte("Golang"))) // true
    fmt.Println(bytes.HasSuffix([]byte("Golang"), []byte("lang")))   // true
    fmt.Println(bytes.HasSuffix([]byte("Golang"), []byte{}))         // true
    fmt.Println(bytes.HasSuffix([]byte("Golang"), nil))              // true
    fmt.Println(bytes.HasSuffix(nil, nil))                           // true

    //定位查找 Index
    // 定位查找(string)
    fmt.Println(strings.Index("Learn Golang, Go!", "Go"))          // 6
    fmt.Println(strings.Index("Learn Golang, Go!", ""))            // 0
    fmt.Println(strings.Index("Learn Golang, Go!", "Java"))        // -1
    fmt.Println(strings.IndexAny("Learn Golang, Go!", "Java"))     // 2
    fmt.Println(strings.IndexRune("Learn Golang, Go!", rune('a'))) // 2

    // 定位查找([]byte)
    fmt.Println(bytes.Index([]byte("Learn Golang, Go!"), []byte("Go")))   // 6
    fmt.Println(bytes.Index([]byte("Learn Golang, Go!"), nil))            // 0
    fmt.Println(bytes.Index([]byte("Learn Golang, Go!"), []byte("Java"))) // -1
    fmt.Println(bytes.IndexAny([]byte("Learn Golang, Go!"), "Java"))      // 2
    fmt.Println(bytes.IndexRune([]byte("Learn Golang, Go!"), rune('a')))  // 2

    // 反向定位查找(string)
    fmt.Println(strings.LastIndex("Learn Golang, Go!", "Go"))      // 14
    fmt.Println(strings.LastIndex("Learn Golang, Go!", ""))        // 17
    fmt.Println(strings.LastIndex("Learn Golang, Go!", "Java"))    // -1
    fmt.Println(strings.LastIndexAny("Learn Golang, Go!", "Java")) // 9

    // 反向定位查找([]byte)
    fmt.Println(bytes.LastIndex([]byte("Learn Golang, Go!"), []byte("Go")))   // 14
    fmt.Println(bytes.LastIndex([]byte("Learn Golang, Go!"), nil))            // 17
    fmt.Println(bytes.LastIndex([]byte("Learn Golang, Go!"), []byte("Java"))) // -1
    fmt.Println(bytes.LastIndexAny([]byte("Learn Golang, Go!"), "Java"))

    //替换
    // 替换(string)
    fmt.Println(strings.Replace("I love java, java, java!!", "java", "go", 1))  // I love go, java, java!!
    fmt.Println(strings.Replace("I love java, java, java!!", "java", "go", 2))  // I love go, go, java!!
    fmt.Println(strings.Replace("I love java, java, java!!", "java", "go", -1)) // I love go, go, go!!
    fmt.Println(strings.Replace("math", "", "go", -1))                          // gomgoagotgohgo
    fmt.Println(strings.ReplaceAll("I love java, java, java!!", "java", "go"))  // I love go, go, go!!
    replacer := strings.NewReplacer("java", "go", "python", "go")
    fmt.Println(replacer.Replace("I love java, python, go!!")) // I love go, go, go!!

    // 替换([]byte)
    fmt.Printf("%s\n", bytes.Replace([]byte("I love java, java, java!!"), []byte("java"), []byte("go"), 1))  // I love go, java, java!!
    fmt.Printf("%s\n", bytes.Replace([]byte("I love java, java, java!!"), []byte("java"), []byte("go"), 2))  // I love go, go, java!!
    fmt.Printf("%s\n", bytes.Replace([]byte("I love java, java, java!!"), []byte("java"), []byte("go"), -1)) // I love go, go, go!!
    fmt.Printf("%s\n", bytes.Replace([]byte("math"), nil, []byte("go"), -1))                                 // gomgoagotgohgo
    fmt.Printf("%s\n", bytes.ReplaceAll([]byte("I love java, java, java!!"), []byte("java"), []byte("go")))  // I love go, go, go!!


}
