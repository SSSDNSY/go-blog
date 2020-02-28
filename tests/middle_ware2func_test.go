package test

//func SingleHost2(handler http.Handler, allowedHost string) http.Handler {
//	fn := func(w http.ResponseWriter, r *http.Request) {
//		if r.Host == allowedHost {
//			handler.ServeHTTP(w, r)
//		} else {
//			w.WriteHeader(403)
//		}
//	}
//	return http.HandleFunc(fn)
//}
//
//func TestFuncServe(t *testing.T) {
//
//}
