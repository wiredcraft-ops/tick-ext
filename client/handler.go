package client

import (
	"fmt"
	"net/http"

	"github.com/Wiredcraft/tick-ext"
	"github.com/Wiredcraft/tick-ext/client/license"
	"github.com/Wiredcraft/tick-ext/common"
	"github.com/abbot/go-http-auth"
	"gopkg.in/square/go-jose.v1/json"
)

var (
	htpasswdPath = "/var/lib/tick-ext/client/.htpasswd"
)

// web ui (auth)
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if basicAuth(w, r) {

		fmt.Fprint(w, "/")
	}
}

// FIXME: 17/1/6
// bcrypt test fail
// panic when htpasswd error
func basicAuth(w http.ResponseWriter, r *http.Request) bool {

	ba := auth.NewBasicAuthenticator(fmt.Sprintf("%s / %s", tick.Name, tick.Version), auth.HtpasswdFileProvider(htpasswdPath))
	if ba.CheckAuth(r) == "" {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return false
	}

	return true
}

// api
func LicenseHandler(w http.ResponseWriter, r *http.Request) {

	msg := common.Msg{}
	msgInJson := []byte("")
	if r.Method == "POST" {
		// validate license
		l := new(license.License)
		err := json.NewDecoder(r.Body).Decode(l)
		if err != nil {
			msg.Success = false
			msg.Message = err.Error()
		} else {
			ret, err := license.Validate(l.Key, common.WclSolutionsApiUrl+"/license")
			if err != nil {
				msg.Success = false
				msg.Message = err.Error()
			}
			if ret {
				msg.Success = true
				msg.Message = "valid key"
			}
		}

		msgInJson, _ = json.Marshal(msg)
		fmt.Fprintf(w, "%s", msgInJson)
	}
}
