package userCron

import (
	"github.com/robfig/cron"
	"github.com/tmaio/go-gin-example/models"
	"log"
	"time"
)

func DoUserJob()  {
	c := cron.New()

	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.Start()

	t := time.NewTimer(time.Second * 10)
	for true {
		select {
		case <-t.C:
			t.Reset(time.Second * 10)
		}
	}

/*	select {

	}*/
}