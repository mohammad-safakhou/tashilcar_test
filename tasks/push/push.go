package push

import (
	"encoding/json"
	"tashil/models"
	"tashil/tasks"
	"tashil/utils"
)

type PusherHandler struct {
}

func EndpointPusher(payload models.RegisterEndpoint) error {
	req, _ := json.Marshal(payload.Endpoint)
	err, _ := tasks.SchedulerTaskHandler(payload.Scheduling.Repeater, req)
	if err != nil {
		return err
	}
	return nil
}

func HttpCall() error {
	return nil
}
