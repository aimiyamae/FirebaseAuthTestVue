package main

import (
	"fmt"
	"html/template"
	"io/ioutil"

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
		// log.Printf("Verified ID token: %v\n", token)
		//firebaseのUID
		log.Printf("%v\n", token.UID)
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

//wikiのデータ構造
type Page struct {
	Title string //タイトル
	Body  []byte //タイトルの中身
}

//パスのアドレスを設定して文字の長さを定数として持つ
const lenPath = len("/view/")

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, err := loadPage(title)
	if err != nil {
		// editのHandlerのURLに飛ばすことで送信ページに飛ばすことができる
		http.Redirect(w, r, "edit/"+title, http.StatusNotFound)
	}
	//view.htmlをgo言語のtemplateパッケージで読み取ってくる
	renderTmplate(w, "view", p)
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	//send.htmlをgo言語のtemplateパッケージで読み取ってくる
	renderTmplate(w, "send", p)
}

func newsHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	//send.htmlをgo言語のtemplateパッケージで読み取ってくる
	renderTmplate(w, "news", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTmplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//tmpl.htmlの中にTitleやBodyを入れれるようにする
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//テキストファイルの保存メソッド
func (p *Page) save() error {
	//タイトルの名前でテキストファイルを作成して保存します。
	filename := p.Title + ".txt"
	//0600は、テキストデータを書き込んだり読み込んだりする権限を設定しています。
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//titleからファイル名を読み込んで新しいPageのポインタを返す
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	//errに値が入ったらエラーとしてbodyの値をnilにして返す
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	//resourcesディレクトリ以下の静的ファイル(ここでいうcss/index.css)を探して返す
	// fmt.Println("The Server runs with https://localhost:8080/")
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("/resources/"))))

	//データの保存と出力
	p1 := &Page{Title: "testPage", Body: []byte("This is a sample page.")}
	p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))
	//http
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/send/", sendHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/news/", newsHandler)
	//コメントアウトを外すとページ遷移　データ繊維できなくなる
	// http.ListenAndServe(":8080", nil)

	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

	r := mux.NewRouter()
	r.HandleFunc("/public", public)
	// r.HandleFunc("/private", private)
	r.HandleFunc("/private", authMiddleware(private))

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
