package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var GoogleClientID string = "914880682379-jd1abklbbmr1p7mh53qcteaatbg1ekne.apps.googleusercontent.com"
var GoogleClientSecret string = "ZSjKZ30jjsvBs0dJZnnir1Xh"
var disableHttp2 *bool = flag.Bool("disable-http2", false, "Disable http2 for this repro")

func main() {
	flag.Parse()

	googleProvider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
	if err != nil {
		panic(err)
	}
	cfg := &oauth2.Config{
		ClientID:     GoogleClientID,
		ClientSecret: GoogleClientSecret,
		Scopes:       []string{oidc.ScopeOpenID},
		Endpoint:     googleProvider.Endpoint(),
		RedirectURL:  "http://localhost:7000/callback",
	}
	server := Server{cfg}

	router := mux.NewRouter()
	router.Methods("GET").Path("/").HandlerFunc(server.sendToLogin)
	router.Methods("GET").Path("/callback").HandlerFunc(server.googleCallback)

	http.ListenAndServe(":7000", router)
}

type Server struct {
	oauthConfig *oauth2.Config
}

func (s *Server) sendToLogin(resp http.ResponseWriter, req *http.Request) {
	url := s.oauthConfig.AuthCodeURL("protect", oauth2.AccessTypeOnline)
	http.Redirect(resp, req, url, http.StatusFound)
}

func (s *Server) googleCallback(resp http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	authenticationCode := req.URL.Query().Get("code")

	token, err := s.oauthConfig.Exchange(ctx, authenticationCode)
	if err != nil {
		panic(err)
	}
	resp.Write([]byte(token.Extra("id_token").(string)))
}
