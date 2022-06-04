package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/zh-five/xdaemon"
	"net/http"
	"os/exec"
)

func main() {
	logFile := "daemon.log"
	d := xdaemon.NewDaemon(logFile)
	d.MaxCount = 2
	d.Run()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		command := "sensors"
		result := RunShell(command)
		c.String(http.StatusOK, result)
	})
	err := r.Run(":80")
	if err != nil {
		slog.Errorf("gin error %v", err)
		return
	}
}

func RunShell(command string) string {
	cmd := exec.Command(command)
	var stdout = bytes.Buffer{}
	var stderr = bytes.Buffer{}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		slog.Error(err)
	}
	slog.Infof(">>>>> stdout:%v\nstderr:%v", stdout.String(), stderr.String())
	return stdout.String()
}
