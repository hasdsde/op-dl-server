package define

import "hasdsd.cn/op-dl-server/model"

type VersionDetails struct {
	model.Version
	model.VersionEvent
	model.VersionTag
	model.Tag
}
type DataWithTotal struct {
	Data  interface{}
	Total int64
}
