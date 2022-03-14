package main

import (
	"github.com/goproxy/goproxy"
	"net/http"
	"os"
)

func main() {
	http.ListenAndServe("localhost:8080", &goproxy.Goproxy{
		GoBinEnv: append(
			os.Environ(),
			"GOPROXY=https://goproxy.cn,direct", // Use Goproxy.cn as the upstream proxy
			"GOPRIVATE=git.example.com",         // Solve the problem of pulling private modules
		),
		ProxiedSUMDBs: []string{
			"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org", // Proxy the default checksum database
		},
	})
}
