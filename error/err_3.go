package main

//哨兵错误处理策略
//反模式
//造成严重的隐式耦合：
//错误值构造方不经意间的一次错误描述字符串的改动，都会造成错误处理方处理行为的改变，并且这种通过字符串比较的方式对错误值进行检视的观感也很差
func main() {
    data, err := b.Peek(1)
    if err != nil {
        switch err.Error {
        case "bufio: negative count":
            return
        case "bufio: buffer full":
            return
        default:
            return
        }
    }
}
