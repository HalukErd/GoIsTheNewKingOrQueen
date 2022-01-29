Check Below
```go
http.Handle("/", http.FileServer(http.Dir(".")))
```

Check FileServer method
```go
func FileServer(root FileSystem) Handler {
	return &fileHandler{root}
}
```

FileSystem is an interface
```go
type FileSystem interface {
	Open(name string) (File, error)
}
```

Check type Dir
```go
type Dir string
```
```go
func (d Dir) Open(name string) (File, error) {
}
```

So you can use type conversion from string to Dir 
and you can use Dir type as FileSystem interface