package main

import (
	"fmt"
	"github.com/aeolee/cron"
)

func main(){
	fmt.Println("Starting...")

	i := 0
	spec := "*/1 * * * * *"
	c := cron.New(cron.WithSeconds())
	c.Start()
	c.AddFunc(spec,func(){
		i++
		fmt.Printf("定时器第 %d 次执行。\n",i)
	})
	select{}
}