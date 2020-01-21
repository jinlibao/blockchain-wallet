package main

import (
	"net/http"

	"github.com/jinlibao/blockchain-wallet/server/cli"
)

func getRespHandler_TemplateFunction(cc *cli.CLI) HandlerFunc {
	return func(www http.ResponseWriter, req *http.Request) {
		if !ValidateAcctPin(cc, www, req) {
			return
		}
		// xyzzy - todo
	}
}
