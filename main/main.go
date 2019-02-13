package main

import (
	"fmt"
	"log"
	"net/http" //http プロトコルを利用してくれるパッケージ
	"os"
	"strings"

	"golang.org/x/net/context"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Firebase SDK のセットアップ
		opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		// クライアントから送られてきた JWT 取得
		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT の検証
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// JWT が無効なら Handler に進まず別処理
			fmt.Printf("error verifying ID token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error verifying ID token\n"))
			return
		}
		log.Printf("Verified ID token: %v\n", token)
		next.ServeHTTP(w, r)
	}
}

// type Ping struct {
// 	Status int    `json:"status"`
// 	Result string `json:"result"`
// }

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	ping := Ping{http.StatusOK, "ok"}
// 	res, _ := json.Marshal(ping)

// 	w.Header().Set("Content-Type", "application/json")

// 	w.Header().Set("Access-Control-Allow-Origin", "*")

// 	w.Write(res)
// }

//httpパッケージ
func handler(writer http.ResponseWriter, request *http.Request) {
	//Hello Worldと言う文字列をwriterに入れて渡す
	//urlのアドレスから文字を取得して表示する http://localhost:8080/aiai⇨Hello!!,aiai
	fmt.Fprintf(writer, "Hello!!,%s", request.URL.Path[1:])
}

func public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello public!\n"))
}

func private(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello private!\n"))
}

func main() {

	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

	r := mux.NewRouter()
	r.HandleFunc("/public", public)
	// r.HandleFunc("/private", private)
	r.HandleFunc("/private", authMiddleware(private))
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
