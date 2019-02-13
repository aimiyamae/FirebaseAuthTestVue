package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

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
	fmt.Println("The Server runs with https://localhost:8070/")
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("/resources/"))))

	//データの保存と出力
	p1 := &Page{Title: "testPage", Body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
	//http
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/send/", sendHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/news/", newsHandler)
	http.ListenAndServe(":8070", nil)
}
