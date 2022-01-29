Check Some error handle approaches
line:13
```go
log.Fatal(http.ListenAndServe(":8080", nil))
```

line:28
```go
if err != nil {
		http.Error(w, "File Not Found", http.StatusNotFound)
	}
```