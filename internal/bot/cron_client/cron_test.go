package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

type Job1 struct {
}

func (t Job1) Run() {
	fmt.Println(time.Now(), "I'm Job1")
}

type Job2 struct {
}

func (t Job2) Run() {
	fmt.Println(time.Now(), "I'm Job2")
}

func TestCorn(t *testing.T) {
	c := cron.New(cron.WithSeconds())

	EntryID, err := c.AddJob("*/5 * * * * *", Job1{})
	fmt.Println(time.Now(), EntryID, err)

	EntryID, err = c.AddJob("*/10 * * * * *", Job2{})
	fmt.Println(time.Now(), EntryID, err)

	c.Start()
	time.Sleep(time.Second * 15)
}
