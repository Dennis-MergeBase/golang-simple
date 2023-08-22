package API

import (
	"fmt"
	"net/http"
)

func UpdateEntity(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Fprintf(w,"Edit Entity")

}
