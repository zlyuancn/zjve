/*
-------------------------------------------------
   Author :       zlyuan
   date：         2019/12/20
   Description :
-------------------------------------------------
*/

package zjve

import (
    "encoding/json"
    "fmt"
    "strings"
)

type Unmarshaler func(data []byte, v interface{}) error

var DefaultUnmarshaler = json.Unmarshal

type JsonValueExtractor struct {
    data        []byte
    mm          map[string]interface{}
    unmarshaler Unmarshaler
}

// 创建一个提取器
func New(data []byte) (*JsonValueExtractor, error) {
    return NewWithUnmarshaler(data, DefaultUnmarshaler)
}

// 创建一个提取器
func NewText(text string) (*JsonValueExtractor, error) {
    return NewWithUnmarshaler([]byte(text), DefaultUnmarshaler)
}

// 创建一个提取器并且使用指定的解码函数
func NewWithUnmarshaler(data []byte, unmarshaler Unmarshaler) (*JsonValueExtractor, error) {
    if unmarshaler == nil {
        unmarshaler = json.Unmarshal
    }

    m := make(map[string]interface{})
    if err := unmarshaler(data, &m); err != nil {
        return nil, err
    }
    return &JsonValueExtractor{
        data:        data,
        mm:          m,
        unmarshaler: unmarshaler,
    }, nil
}

// 获取数据
func (m *JsonValueExtractor) Data() []byte {
    return m.data
}

// 将数据绑定到任何类型的值或指针
func (m *JsonValueExtractor) Unmarshal(outPtr interface{}) error {
    return m.unmarshaler(m.data, outPtr)
}

// 读取
func (m *JsonValueExtractor) Get(path string) (interface{}, error) {
    return m.GetOfSep(path, ".")
}

// 读取并自定义分隔符
func (m *JsonValueExtractor) GetOfSep(path, sep string) (interface{}, error) {
    if path == "" {
        return nil, nil
    }

    ks := strings.Split(path, sep)
    var value interface{} = m.mm
    var ok bool
    var mm (map[string]interface{})
    for i, k := range ks {
        if mm, ok = value.(map[string]interface{}); !ok {
            return nil, fmt.Errorf("路径 %s 中 %s 是一个具体值", path, strings.Join(ks[:i], sep))
        }

        if value, ok = mm[k]; !ok {
            return nil, fmt.Errorf("路径 %s 中没有找到 %s", path, strings.Join(ks[:i+1], sep))
        }
    }
    return value, nil
}
