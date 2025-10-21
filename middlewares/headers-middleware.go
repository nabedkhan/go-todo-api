package middlewares

import "net/http"

// Middlewares
// func HeadersMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")

// 		next.ServeHTTP(w, r)
// 	})

// }

func HeadersMiddleware(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Standard JSON Response Header
		w.Header().Set("Content-Type", "application/json")

		// CORS Headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, PUT OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next(w, r)
	}

	return http.HandlerFunc(handler)

}
