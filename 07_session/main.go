package main

import "go_web/07_session/utils"

var globalSessions *utils.Manager

func init() {
	globalSessions, _ = utils.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}
func main() {

}
