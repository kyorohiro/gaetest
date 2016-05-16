package hello

import (
	"encoding/json"
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
		ctx := appengine.NewContext(r)
		data := GetParam(r)
		user := gaeuser.NewUser(ctx, data["name"].(string))
		err := user.Regist(ctx, data["pass"].(string), data["mail"].(string))
		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}
		//
		Response(w, map[string]interface{}{"r": "ok", "s": "good"})
	})

	http.HandleFunc("/user/updateMail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!")
		ctx := appengine.NewContext(r)
		data := GetParam(r)
		user := gaeuser.NewUser(ctx, data["name"].(string))
		err := user.UpdateMail(ctx, data["mail"].(string))
		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}
		//
		Response(w, map[string]interface{}{"r": "ok", "s": "good"})
	})

	http.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!")
		ctx := appengine.NewContext(r)
		data := GetParam(r)
		user := gaeuser.NewUser(ctx, data["name"].(string))
		err := user.PullFromDB(ctx)
		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}
		//
		Response(w, map[string]interface{}{
			"r": "ok", "s": "good", //
			"name":     user.GaeObject.UserName, //
			"created":  user.GaeObject.Created,  //
			"logined":  user.GaeObject.Logined,  //
			"loginid":  user.GaeObject.LoginId,  //
			"mail":     user.GaeObject.Mail,     //
			"passHash": user.GaeObject.PassHash, //
			"meicon":   user.GaeObject.MeIcon,   //
		})
	})

	http.HandleFunc("/user/mail/getUser", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!")
		ctx := appengine.NewContext(r)
		data := GetParam(r)

		mail := gaeuser.NewMail(ctx, data["mail"].(string))
		err := mail.PullFromDB(ctx)
		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}
		user, err1 := mail.GetUser(ctx)
		if err1 != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err1.Error()})
			return
		}

		//
		Response(w, map[string]interface{}{"r": "ok", "s": "good",
			"mail_mail":    mail.GaeObject.Mail,     //
			"mail_name":    mail.GaeObject.UserName, //
			"user_name":    user.GaeObject.UserName,
			"user_created": user.GaeObject.Created,
			"user_logined": user.GaeObject.Logined,
			"user_loginid": user.GaeObject.LoginId,
		})
	})
}

func Response(w http.ResponseWriter, v map[string]interface{}) {
	b, _ := json.Marshal(v)
	fmt.Fprintln(w, string(b))
}

func GetParam(r *http.Request) map[string]interface{} {
	var data map[string]interface{}
	json.NewDecoder(r.Body).Decode(&data)
	return data
}
