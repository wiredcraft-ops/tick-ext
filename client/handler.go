package client

import (
	"fmt"
	"github.com/Wiredcraft/tick-ext"
	"github.com/abbot/go-http-auth"
	"net/http"
)

var (
	unauthorized = "401 Unauthorized"
	htpasswdPath = "/var/lib/tick-ext/client/.htpasswd"
)

// auth
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
		http.Error(w, unauthorized, http.StatusUnauthorized)
		return false
	}

	return true
}

// api
func LicenseHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		// validate license
	}
}
