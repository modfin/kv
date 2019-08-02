# KV

A small wrapper around a key value in order to help with parsing of headers and such



```go 

kv1 := kv.New("a", 123.34)
kv2 := kv.New("b", "123.34")
kv3 := kv.New("a", []int{1,2,3,4})

f1, ok := kv1.Float()
f2 := kv1.FloatOr(-1.1)

str1, ok := kv2.String()
str2 := kv2.StringOr("default")

s1, ok := kv3.IntSlice()
s2 := kv3.IntSliceOr([]int{1})


```