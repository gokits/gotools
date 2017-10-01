# gotools
tools for golang

## StructCopy
Copy fields which have same name and types between different types of struct. Note that unexported fields are ignored.

```go
type Src struct {
        A int
        B string
        c byte
}

type Dst struct {
        A int
        B byte
        c byte
}

func main(){
        src := &Src{3, "hello", '2'}
        var dst Dst
        gotools.StructCopy(&dst, src)
        fmt.Println(dst)
}

/* output
{3 0 0}
*/
```
