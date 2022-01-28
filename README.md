*http.Handle gets a Handler that has ServeHTTP(ResponseWriter, *Request)

*http.HandleFunc get a Func type with parameter ResponseWriter and Pointer of Request

*http.HandlerFunc is a type that underlies That Func

*Also http.HandlerFunc has ServeHTTP(ResponseWriter, *Request) so can be used as a Handler

*With type conversion HandlerFunc(that Func) we can acquire a Handler