package taskmanager

import (
	"github.com/pivotal-pez/pezdispenser/service/integrations"
	"labix.org/v2/mgo/bson"
)

type (
	//Task - a task object
	Task struct {
		ID         bson.ObjectId          `bson:"_id"`
		Timestamp  int64                  `bson:"timestamp"`
		Expires    int64                  `bson:"expires"`
		Status     string                 `bson:"status"`
		Profile    ProfileType            `bson:"profile"`
		CallerName string                 `bson:"caller_name"`
		MetaData   map[string]interface{} `bson:"metadata"`
	}

	//TaskManager - manages task interactions crud stuff
	TaskManager struct {
		taskCollection integrations.Collection
	}

	//ProfileType - indicator of the purpose of the task to be performed
	ProfileType string
)
