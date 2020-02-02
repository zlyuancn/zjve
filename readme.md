# Json提取器

---

# 此包以后不再更新, 请转到 [github.com/zlyuancn/zjve2](github.com/zlyuancn/zjve2)

# 获得
`go get -u github.com/zlyuancn/zjve`

# 实例
```
func main() {
    text := `{"a":"a1","b":{"x":"x1","y":{"1":"1"}}}`
    fmt.Println(text)
    be, err := zjve.NewText([]byte(text))
    if err != nil {
        panic(err)
    }

    paths := []struct {
        path string
        ok   bool
    }{
        {"a", true},
        {"a.b", false},
        {"b", true},
        {"b.x", true},
        {"b.x.1", false},
        {"b.y", true},
        {"b.y.1", true},
        {"b.y.1.2", false},
        {"b.y.2", false},
        {"c", false},
        {"c.x", false},
        {"c.y", false},
        {"c.y.1", false},
    }
    var v interface{}
    for _, p := range paths {
        if v, err = be.Get(p.path); err != nil {
            if p.ok {
                fmt.Printf("err [%t]%s:\t\t它应该是成功的,但是出现了错误:%s\n", p.ok, p.path, err)
            } else {
                fmt.Printf("ok [%t]%s:\t\t%v\n", p.ok, p.path, err)
            }
        } else {
            if !p.ok {
                fmt.Printf("err [%t]%s:\t\t它应该是失败的的,但是收到了值:%v\n", p.ok, p.path, v)
            } else {
                fmt.Printf("ok [%t]%s:\t\t%v\n", p.ok, p.path, v)
            }
        }
    }
}
```
