package dist_store


type DistStorer interface {
	Write(key, value string) error
	Read(key string) (string, error)
	Dump()
}