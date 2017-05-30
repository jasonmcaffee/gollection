# gollection
gollection provides commonly needed high order functions for iterating, filtering, reducing, etc slices in go.

## Usage

### Filter
The filter function allows you to filter a given slice utilizing a bool function which indicates the item in the slice should be included in the result.

```go
g.New([]string{"a", "b"}).
    Filter(func(s string) bool {return s == "b"}).
    Each(func(s string){
        fmt.Println(s)
    })
//prints b
```