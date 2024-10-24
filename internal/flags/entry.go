package flags

type Entry struct {
	Flag  string
	Value []string
}

func NewEntry(flag string, value []string) *Entry {
	return &Entry{
		Flag:  flag,
		Value: value,
	}
}
