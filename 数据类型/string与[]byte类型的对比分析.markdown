# Golang string和[]byte的对比
## string类型
在go语言中string类型是utf8类型的只可读字符切片
### string结构
```
    type stringStruct struct {
	    str unsafe.Pointer
	    len int
    }
```
可以看到string结构体由两个字段组成，指向某个数组首地址的指针和字符串的长度len
