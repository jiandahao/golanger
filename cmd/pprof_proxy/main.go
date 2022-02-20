package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	args := append([]string{"tool", "pprof"}, os.Args[1:]...)

	tmpdir, ok := os.LookupEnv("PPROF_TMPDIR")
	if !ok || tmpdir == "" {
		os.Setenv("PPROF_TMPDIR", "/tmp/pprof")
		tmpdir = "/tmp/pprof"
	}

	cmd := exec.Command("go", args...)
	stdout := &customStdout{}
	cmd.Stdout = stdout
	cmd.Stderr = stdout

	if err := cmd.Run(); err != nil {
		fmt.Println(stdout.String())
		panic(err)
	}

	remoteURL := "http://127.0.0.1:8086"

	go func() {
		// RUN go tool pprof -http=:8086 -no_browser ${PPROF_TMPDIR}
		entires, err := os.ReadDir(tmpdir)
		if err != nil {
			panic(err)
		}

		var profileFile string
		for _, entry := range entires {
			if !entry.IsDir() {
				profileFile = filepath.Join(tmpdir, entry.Name())
				break
			}
		}

		if profileFile == "" {
			panic("profile file not found")
		}

		cmd := exec.Command("go", "tool", "pprof", "-http=127.0.0.1:8086", "-no_browser", profileFile)
		cmd.Stdout = stdout
		cmd.Stderr = stdout
		if err := cmd.Run(); err != nil {
			fmt.Println(stdout.String())
			panic(err)
		}
	}()

	// RUN proxy
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	remote, err := url.Parse(remoteURL)
	if err != nil {
		panic("invalid pprof web UI addr")
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(r *http.Request) {
		r.URL.Scheme = remote.Scheme
		r.URL.Host = remote.Host
	}

	router.GET("/*any", func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	fmt.Printf("\n[pprof_proxy] Running web UI, visit: <your host addr>:<your host port>\n e.g 192.168.1.12:8085 \n don't use localhost or 127.0.0.1 when running in container \n\n")

	fmt.Println("Listening and serving Web UI on 127.0.0.1:8085")
	if err := router.Run(":8085"); err != nil {
		panic(err)
	}
}

type customStdout struct {
	buf bytes.Buffer
}

func (s *customStdout) Write(p []byte) (n int, err error) {
	return s.buf.Write(p)
}

func (s *customStdout) String() string {
	return s.buf.String()
}
