package _cronUtils

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"sync"
)

var cronInst *cron.Cron

var taskFuncMap = sync.Map{}

func GetCrontabInstance() *cron.Cron {
	if cronInst != nil {
		return cronInst
	}
	cronInst = cron.New()
	cronInst.Start()

	return cronInst
}

func AddTask(name string, schedule string, f func()) (entryID cron.EntryID, err error) {
	if _, ok := getTaskFuncFromMap(name); !ok {
		fmt.Println("Add a new task:", name)

		cInstance := GetCrontabInstance()
		entryID, err = cInstance.AddFunc(schedule, f)
		if err == nil {
			taskFuncMap.Store(name, entryID)
		}

	} else {
		fmt.Println("Don't add same task `" + name + "` repeatedly!")
	}
	return
}

func RemoveTask(name string) {
	if entryID, ok := getTaskFuncFromMap(name); ok {
		fmt.Println("remove task:", name)
		cInstance := GetCrontabInstance()
		cInstance.Remove(entryID)
		deleteTaskFuncFromMap(name)
	}
}

func Stop() {
	cronInst.Stop()
}

func getTaskFuncFromMap(key string) (ret cron.EntryID, ok bool) {
	obj, ok := taskFuncMap.Load(key)

	if ok {
		ret = obj.(cron.EntryID)
	}

	return
}

func deleteTaskFuncFromMap(key string) {
	taskFuncMap.Delete(key)
	return
}
