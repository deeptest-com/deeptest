package _cronUtils

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

var cronInst *cron.Cron

var taskFunc = make(map[string]cron.EntryID)

func GetCrontabInstance() *cron.Cron {
	if cronInst != nil {
		return cronInst
	}
	cronInst = cron.New()
	cronInst.Start()

	return cronInst
}

func AddTask(name string, schedule string, f func()) (entryID cron.EntryID, err error) {
	if _, ok := taskFunc[name]; !ok {
		fmt.Println("Add a new task:", name)

		cInstance := GetCrontabInstance()
		entryID, err = cInstance.AddFunc(schedule, f)
		if err != nil {
			taskFunc[name] = entryID
		}

	} else {
		fmt.Println("Don't add same task `" + name + "` repeatedly!")
	}
	return
}

func RemoveTask(name string) {
	if entryID, ok := taskFunc[name]; ok {
		fmt.Println("remove task:", name)
		cInstance := GetCrontabInstance()
		cInstance.Remove(entryID)
	}
}

func Stop() {
	cronInst.Stop()
}
