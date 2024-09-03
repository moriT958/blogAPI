package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/moriT958/go-api/models"
)

// GET /helloのハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	//　wは、http.ResponseWriterで定義されたインターフェース型。reqをwに埋め込む形で使う。
	// reqは、http.Request構造体のポインタ。

	// WriteStringはio.Writerに文字列を埋め込める。
	io.WriteString(w, "Hello, World!!\n")
}

// POST /articleのハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	// 1. 受け取ったJSONをバイトとして読み込む処理
	// リクエストヘッダからバイトの長さを得る。
	length, err := strconv.Atoi(req.Header.Get("Content-Length")) // Atoiで文字列を数字に変換
	if err != nil {
		http.Error(w, "failed to get Content-Length\n", http.StatusBadRequest)
		return
	}
	reqBuffer := make([]byte, length) // 得た長さに合わせてバイトスライスを作成

	// リクエストボディをバイトに読み込む。
	if _, err := req.Body.Read(reqBuffer); !errors.Is(err, io.EOF) {
		// errがEOFでない(最後まで読み込めなかった)時の処理。
		http.Error(w, "failed to get request body\n", http.StatusBadRequest)
		return
	}
	defer req.Body.Close() // 処理の最後に閉じる

	// 2. バイトからGoの構造体に変換する処理
	var reqArticle models.Article                                  // 構造体を初期化しておく
	if err := json.Unmarshal(reqBuffer, &reqArticle); err != nil { // Unmarshalでjsonのバイトを構造体に埋め込む。
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
		return
	}

	article := reqArticle
	jsonData, err := json.Marshal(article) // json.Marshalで構造体をJSONに変換。
	if err != nil {
		http.Error(w, "Failed to encode to json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// GET /article/list?page= のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// リクエストからURLのクエリ部を得る。
	queryMap := req.URL.Query() // Queryはマップ(url.Values)を返す

	var page int // リストデータをページに分けて返すようにする。
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		// page=1&page=2の場合、"page"に対応するマップの値は[1,2]をして返ってくるため、len(p)と書ける。
		var err error

		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMsg := fmt.Sprintf("Faild to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// GET /article/{id}のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	// mux.Varsはリクエストのルートの値をマップで返す。
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("Faild to encode json (articleID %d)\n", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
	}

	w.Write(jsonData)
}

// POST /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Faild to encode json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// POST /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment1 := models.Comment1
	jsonData, err := json.Marshal(comment1)
	if err != nil {
		http.Error(w, "Faild to encode json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
