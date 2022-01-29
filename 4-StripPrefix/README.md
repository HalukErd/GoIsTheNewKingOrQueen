Check Below
```go
http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
```

