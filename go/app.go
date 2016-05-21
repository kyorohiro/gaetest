package hello

import (
	"encoding/json"
	"fmt"
	"net/http"
	"umiuni2d_backend/twitter"
	"umiuni2d_backend/user"

	"google.golang.org/appengine/log"

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
		userManager := gaeuser.NewUserManager("testuser", "testloginid")
		user, err := userManager.RegistUser(ctx, data["name"].(string), data["pass"].(string), data["mail"].(string))
		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}
		//
		Response(w, map[string]interface{}{"r": "ok", "s": "good", "p": user.GetUserName()})
	})

	http.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		data := GetParam(r)
		userMana := gaeuser.NewUserManager("testuser", "testloginid")
		user, err := userMana.FindUserFromUserName(ctx, data["name"].(string))
		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}
		//
		Response(w, map[string]interface{}{
			"r": "ok", "s": "good", //
			"name":     user.GetUserName(), //
			"created":  user.GetCreated(),  //
			"logined":  user.GetLogined(),  //
			"mail":     user.GetMail(),     //
			"passHash": user.GetPassHash(), //
			"meicon":   user.GetMeIcon(),   //
		})
	})

	http.HandleFunc("/user/check", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		data := GetParam(r)
		log.Infof(ctx, "###########---------------/user/check(1a)")
		userMana := gaeuser.NewUserManager("testuser", "testloginid")
		log.Infof(ctx, "###########---------------/user/check(1b)")

		isLoginA, _, _ := userMana.CheckLoginId(ctx, data["loginId"].(string), r.RemoteAddr, r.UserAgent())
		log.Infof(ctx, "###########---------------/user/check(1c)")

		isLoginB, _, _ := userMana.CheckLoginId(ctx, data["loginId"].(string), r.RemoteAddr, r.UserAgent())
		log.Infof(ctx, "###########---------------/user/check(1d)")

		isLoginC, _, _ := userMana.CheckLoginId(ctx, data["loginId"].(string), r.RemoteAddr, r.UserAgent())
		log.Infof(ctx, "###########---------------/user/check(2)")

		//
		Response(w, map[string]interface{}{
			"r": "ok", "s": "good", //
			"loginA": isLoginA, //
			"loginB": isLoginB, //
			"loginC": isLoginC, //
		})
	})
	http.HandleFunc("/user/updateMail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!!")
		ctx := appengine.NewContext(r)
		data := GetParam(r)
		userMana := gaeuser.NewUserManager("testuser", "testloginid")
		userObj, err := userMana.FindUserFromUserName(ctx, data["name"].(string)) //UpdateMail(ctx, data["name"].(string), data["mail"].(string))

		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}

		userObj.SetMail(data["mail"].(string))
		err = userObj.PushToDB(ctx)
		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}
		//
		Response(w, map[string]interface{}{"r": "ok", "s": "good"})
	})

	http.HandleFunc("/user/mail/getUser", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, "Hello World!!")
		ctx := appengine.NewContext(r)
		data := GetParam(r)

		userMana := gaeuser.NewUserManager("testuser", "testloginid")
		user, err1 := userMana.FindUserFromMail(ctx, data["mail"].(string))
		if err1 != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err1.Error()})
			return
		}

		//
		Response(w, map[string]interface{}{"r": "ok", "s": "good",
			"mail_mail":    user.GetMail(), //
			"user_name":    user.GetUserName(),
			"user_created": user.GetCreated(),
			"user_logined": user.GetLogined(),
		})
	})

	http.HandleFunc("/user/login", func(w http.ResponseWriter, r *http.Request) {
		r.UserAgent()
		//r.RemoteAddr
		ctx := appengine.NewContext(r)
		data := GetParam(r)

		userMana := gaeuser.NewUserManager("testuser", "testloginid")
		log.Infof(ctx, "##login ")
		loginId, user, err := userMana.LoginUser(ctx, data["name"].(string), data["pass"].(string), r.RemoteAddr, r.UserAgent())
		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}

		//
		Response(w, map[string]interface{}{ //
			"r": "ok", "s": "good", //
			"loginId":   loginId.GetLoginId(), //
			"user_name": user.GetUserName(),
			"dev":       r.UserAgent(),
		})
	})
	http.HandleFunc("/user/logout", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		data := GetParam(r)

		userMana := gaeuser.NewUserManager("testuser", "testloginid")
		log.Infof(ctx, "##logout "+data["loginId"].(string))
		err := userMana.LogoutUser(ctx, data["loginId"].(string), r.RemoteAddr, r.UserAgent())
		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}

		//
		Response(w, map[string]interface{}{ //
			"r": "ok", "s": "good", //
		})
	})

	http.HandleFunc("/user/delete", func(w http.ResponseWriter, r *http.Request) {

		ctx := appengine.NewContext(r)
		data := GetParam(r)

		userMana := gaeuser.NewUserManager("testuser", "testloginid")
		err := userMana.DeleteUser(ctx, data["name"].(string), data["pass"].(string))
		if err != nil {
			Response(w, map[string]interface{}{"r": "ng", "s": err.Error()})
			return
		}

		//
		Response(w, map[string]interface{}{ //
			"r": "ok", "s": "good", //
		})
	})

	http.HandleFunc("/twitter", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		///		data := GetParam(r)

		twitterObj := twitter.NewTwitter(consumerKey, consumerSecret, accessToken, accessTokenSecret, "http://127.0.0.1:8080/oauth")
		url, _, err := twitterObj.SendRequestToken(ctx)
		if err != nil {
			Response(w, map[string]interface{}{ //
				"r": "ng", "s": "good", "dev": err.Error(), //
			})
			return
		}
		http.Redirect(w, r, url, http.StatusFound)

	})

	http.HandleFunc("/twitter/oauth", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		log.Infof(ctx, "=======OKK-Z----->")
		twitterObj := twitter.NewTwitter(consumerKey, consumerSecret, accessToken, accessTokenSecret, callback)
		_, rt, err := twitterObj.OnCallbackSendRequestToken(ctx, r.URL)
		if err != nil {
			Response(w, map[string]interface{}{ //
				"r": "ng", "s": "good", "dev": err.Error(), //
			})
			return
		}
		userMana := gaeuser.NewUserManager("testuser", "testloginid")
		userMana.RegistUserFromTwitter(ctx, rt[twitter.ScreenName], rt[twitter.UserID], rt[twitter.OAuthToken], rt[twitter.OAuthTokenSecret])
		userMana.LoginUserFromTwitter(ctx, rt[twitter.ScreenName], rt[twitter.UserID], rt[twitter.OAuthToken], rt[twitter.OAuthTokenSecret],
			r.RemoteAddr, r.UserAgent())

		Response(w, map[string]interface{}{ //
			"r": "ok", "s": "good",
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
