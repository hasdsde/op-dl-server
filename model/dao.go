package model

type VersionDetails struct {
	Version
	VersionEvent
	VersionTag
	Tag
}
type DataWithTotal struct {
	Data  interface{}
	Total int64
}
