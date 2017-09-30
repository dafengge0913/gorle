# gorle
Run-length encoding
# usage
```go
func main() {
	data := []byte("111111112344444444")
	fmt.Printf("%v \n", data)
	fmt.Printf("%v \n", gorle.Encode(data))
	fmt.Printf("%v \n", gorle.Decode(gorle.Encode(data)))
}
```
