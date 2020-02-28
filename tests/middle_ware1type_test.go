package test

//
//type SingleHost struct {
//	handler     http.Handler
//	allowedHost string
//}
//
//func (this *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	println(r.Host)
//	if r.Host == this.allowedHost {
//		this.handler.ServeHTTP(w, r)
//	} else {
//		w.WriteHeader(403)
//	}
//}
//
//func myHandler(w ResponseWriter, r *Request) {
//	w.Write([]byte("hello world!"))
//}
//
//func TestServer(t *testing.T) {
//	single := &SingleHost{
//		handler:     http.HandlerFunc(myHandler),
//		allowedHost: "example.com",
//	}
//	http.ListenAndServe(":8080", single)
//}
