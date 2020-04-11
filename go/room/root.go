package room

import (
	"github.com/planning_poker/go/common/constant"
	"log"
	"net/http"
	"strings"
)

type prototype func(w http.ResponseWriter, r *http.Request)

var postMethodMap = map[string]prototype{
	"":,
	"/join":,

}

vargit putMethodMap = map[string]prototype{
	"":,
	"/participant":,
}

var getMethodMap = map[string]prototype{
	"": region.Read,
}

var methodMap = map[string]map[string]prototype{
	http.MethodPost: postMethodMap,
	http.MethodPut: putMethodMap,
	http.MethodGet: getMethodMap,
}

func Dispatch(w http.ResponseWriter, r *http.Request) {

	log.Printf("daily.Dispatch : %s", r.Method)

	// HttpMethodに該当するものがない場合はエラー
	mMap := methodMap[r.Method]
	if mMap == nil {
		// Methodがない場合はNotAllowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("daily.Dispatch full-path : %s", r.URL.Path)

	// ルートを削除したパスを取得する
	path := strings.Replace(r.URL.Path, constant.RoomRootPath, "", -1)

	log.Printf("daily.Dispatch path : %s", path)

	// 対応するFunctionを取得する
	method := mMap[path]
	if method == nil {
		// Functionがない場合は、NotFound
		w.WriteHeader(http.StatusNotFound)
		return
	}

	method(w, r)
}