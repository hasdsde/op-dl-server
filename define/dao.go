package define

import "hasdsd.cn/op-dl-server/model"

type VersionDetails struct {
	model.Version
	//versionEvent model.VersionEvent
	//versionTag   []model.VersionTag
	//tag          []model.Tag
	TagID     string
	VersionId string
}
type DataWithTotal struct {
	Data  interface{}
	Total int64
}
