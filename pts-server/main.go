package main

import (
	"github.com/astaxie/beego"
	"fmt"
	"time"
	"strings"
	"net/http"
)

type PingController struct {
	beego.Controller
}

func (this *PingController) Get() {
	rl.Logs = append(rl.Logs, fmt.Sprintf("Agent: %s is started at %s.", this.Ctx.Request.RemoteAddr, time.Now().Format("2006-01-02 15:04:05")))
	this.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
}

type PongController struct {
	beego.Controller
}

func (this *PongController) Get(){
	rl.Logs = append(rl.Logs, fmt.Sprintf("Standby Agent: %s is ready at %s.", this.Ctx.Request.RemoteAddr, time.Now().Format("2006-01-02 15:04:05")))
	this.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
}

type SummaryController struct {
	beego.Controller
}

func (this *SummaryController) Get() {
	this.Ctx.WriteString(strings.Join(rl.Logs, "\n"))
}

type RequestLogs struct {
	Logs [] string
}

var rl *RequestLogs

func init() {
	rl = &RequestLogs{
		Logs: make([]string, 0),
	}
}

func main() {
	beego.Router("/ping", &PingController{})
	beego.Router("/pong", &PongController{})
	beego.Router("/", &SummaryController{})
	beego.Run()
}
