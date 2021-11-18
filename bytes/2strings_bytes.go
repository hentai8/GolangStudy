package main

import (
    "bytes"
    "fmt"
    "strings"
)
func Equal(a, b []byte) bool {
    return string(a) == string(b)
}

func Compare(a, b string) int {
    if a == b {
        return 0
    }
    if a < b {
        return -1
    }
    return +1
}


func main() {
    //比较
    var a = []byte{'a', 'b', 'c'}
    //var b = []byte{'a', 'b', 'd'}

    //if a == b { // 错误：invalid operation: a == b
    //    fmt.Println("slice a is equal to slice b")
    //} else {
    //    fmt.Println("slice a is not equal to slice b")
    //}

    if a != nil { // 正确：valid operation
        fmt.Println("slice a is not nil")
    }

    // 但 Go 语言原生支持通过操作符 == 或 != 对 string 类型变量进行等值比较，
    //因此 strings 包未像 bytes 包那样提供了 Equal 函数。
    //而 bytes 包的 Equal 函数的实现也是基于原生字符串类型的等值比较的：

    fmt.Println(bytes.Equal([]byte{'a', 'b', 'c'}, []byte{'a', 'b', 'd'})) // false "abc" != "abd"
    fmt.Println(bytes.Equal([]byte{'a', 'b', 'c'}, []byte{'a', 'b', 'c'})) // true  "abc" == "abc"
    fmt.Println(bytes.Equal([]byte{'a', 'b', 'c'}, []byte{'b', 'a', 'c'})) // false "abc" != "bac"
    fmt.Println(bytes.Equal([]byte{}, nil))

    // EqualFold 不区分大小写 Equal 区分大小写
    fmt.Println(strings.EqualFold("GoLang", "golang"))               // true
    fmt.Println(bytes.Equal([]byte("GoLang"), []byte("Golang")))     // false
    fmt.Println(bytes.EqualFold([]byte("GoLang"), []byte("Golang"))) // true


    a = []byte{'a', 'b', 'c'}
    var b = []byte{'a', 'b', 'd'}
    var c = []byte{} //empty slice
    var d []byte     //nil slice

    fmt.Println(bytes.Compare(a, b))     // -1 a < b
    fmt.Println(bytes.Compare(b, a))     // 1 b < a
    fmt.Println(bytes.Compare(c, d))     // 0
    fmt.Println(bytes.Compare(c, nil))   // 0
    fmt.Println(bytes.Compare(d, nil))   // 0
    fmt.Println(bytes.Compare(nil, nil)) // 0 nil == nil
    //实际应用中，我们很少使用 strings.Compare，更多的是直接使用排序比较操作符对字符串类型变量进行比较。


    a = []byte{'a', 'b', 'c'}
    b = []byte{'a', 'b', 'd'}
    c = []byte{} //empty slice
    d = nil     //nil slice

    fmt.Println(bytes.Compare(a, b))     // -1 a < b
    fmt.Println(bytes.Compare(b, a))     // 1 b < a
    fmt.Println(bytes.Compare(c, d))     // 0
    fmt.Println(bytes.Compare(c, nil))   // 0
    fmt.Println(bytes.Compare(d, nil))   // 0
    fmt.Println(bytes.Compare(nil, nil)) // 0 nil == nil

    // split 分割
    // 空白分割的字符串是我们能遇到的最简单的也是最常见的由特定分隔符分隔的数据，strings 包和 bytes 包中的 Fields 函数可直接用于处理这类数据的分割，我们看下面示例：
    fmt.Printf("%q\n", strings.Fields("go java python"))                         // ["go" "java" "python"]
    fmt.Printf("%q\n", strings.Fields("\tgo  \f \u0085 \u00a0 java \n\rpython")) // ["go" "java" "python"]
    fmt.Printf("%q\n", strings.Fields(" \t \n\r   "))                            // []

    fmt.Printf("%q\n", bytes.Fields([]byte("go java python")))                         // ["go" "java" "python"]
    fmt.Printf("%q\n", bytes.Fields([]byte("\tgo  \f \u0085 \u00a0 java \n\rpython"))) // ["go" "java" "python"]
    fmt.Printf("%q\n", bytes.Fields([]byte(" \t \n\r   ")))                            // []

    //拼接 Concatenate
    s := []string{"I", "love", "Go"}
    fmt.Println(strings.Join(s, " ")) // I love Go
    p := [][]byte{[]byte("I"), []byte("love"), []byte("Go")}
    fmt.Printf("%q\n", bytes.Join(p, []byte(" "))) // "I love Go"

    ss := []string{"I", "love", "Go"}
    var builder strings.Builder
    for i, w := range ss {
        builder.WriteString(w)
        if i != len(ss)-1 {
            builder.WriteString(" ")
        }
    }
    fmt.Printf("%s\n", builder.String()) // I love Go

    bb := [][]byte{[]byte("I"), []byte("love"), []byte("Go")}
    var buf bytes.Buffer
    for i, w := range bb {
        buf.Write(w)
        if i != len(b)-1 {
            buf.WriteString(" ")
        }
    }
    fmt.Printf("%s\n", buf.String()) // I love Go
}

