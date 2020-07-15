package resource

import "time"

const HTTPTimeOut = 5 * time.Second
const RundeckAuthHeaderName = "X-Rundeck-Auth-Token"

type ProjectResquest struct {
	Name *string `json:",omitempty"`
}
