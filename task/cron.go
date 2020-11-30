package task

import (
	"fmt"
	//"qtmd-php/handlers"
	"github.com/robfig/cron/v3"
)

type CronTask struct {
	Name string
}

func (t *CronTask) Run() {
	fmt.Println("cron task")
}

func AddCron() {
	c := cron.New(cron.WithSeconds())
	//spec := "0 */10 9-15 * * 1-5"
	//spec := "*/5 * * * * *"
	//c.AddFunc(spec, func() {
	//	handlers.KzzTest()
	//})

	//yybspec := "* */1 * * * *"
	//c.AddFunc(yybspec, func() {
	//	handlers.KzzYyb()
	//})

	//c.AddJob("*/5 * * * * *", &CronTask{Name: "testcron"})
	c.Start()
	defer c.Stop()
	select {}
}
