package hello

import (
	"fmt"
	"gaeuser"
	"net/http"

	//	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!")
	})

	http.HandleFunc("/user/new", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!")
		data := GetParam(r)
		user := gaeuser.NewUser(appengine.NewContext(r), data["name"].(string))
		err := user.PullFromDB(ctx)
		err = user.Regist(ctx, data["pass"].(string), data["mail"].(string))
		//
		Response(map[string]string{"r": "ok", "s": "good"})
	})
}

func Response(v map[string]interface{}) {
	b, _ := json.Marshal(v)
	fmt.Fprintln(w, string(b))
}

func GetParam(r *http.Request) map[string]interface{} {
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)
	return data
}
