package utils

import "github.com/guonaihong/gout"

// HttpClient 全局 Http Client
var HttpClient *gout.Client

func init() {
	HttpClient = gout.NewWithOpt(gout.WithInsecureSkipVerify())
}
