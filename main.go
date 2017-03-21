package main

import (
	"flag"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var GoogleClientID string = "190037601363-v72t2b5e7g7vaija4p3mqcmtj9gq8oi7.apps.googleusercontent.com"
var GoogleClientSecret string = "K9yGv8j3P2zzak8RVC0xtwOI"
var disableHttp2 *bool = flag.Bool("disable-http2", false, "Disable http2 for this repro")

func main() {
	flag.Parse()

}
