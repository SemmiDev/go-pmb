package registrant

type ID string

func (i ID) String() string {
	return string(i)
}

func (i ID) Empty() bool {
	return string(i) == ""
}
