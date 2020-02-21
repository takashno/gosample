package main

import (
    "fmt"
    "golang.org/x/text/width"
    piyo "github.com/takashno/gosample2/mypkg"
)

func main() {
    // 半角文字を全角文字にするプログラム
    text := "123XYZ"
    fmt.Println("半角:", text)
    fmt.Println("全角:", width.Widen.String(text))
    fmt.Println(piyo.Hoge())
}