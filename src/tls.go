package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
)

func setupTls(r *gin.Engine) {
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Email:      "*****",
		HostPolicy: autocert.HostWhitelist("*****"),
		Cache:      autocert.DirCache("*****"),
	}

	log.Fatal(autotls.RunWithManager(r, &m))
}
