package fetcher

type Filter struct {
	owner    uint8
	others   uint8
	all      uint8
	extended uint8
}

type Wall struct {
	owner_id int64
	domain   string
	offset   uint64
	count    uint8
}
