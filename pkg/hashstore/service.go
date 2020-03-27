package hashstore

type Storer interface {
	Write(key, value string) error
	Read(key string) (string, error)
	Dump() map[string]string
	Healthy() bool
}

//type DumpIterator interface {
//	H
//}