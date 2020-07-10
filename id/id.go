package id

type ID string

var NilID ID = ""

func (id ID) String() string {
	return string(id)
}

func (id ID) Nil() bool {
	return id == NilID
}

func FromString(id string) ID {
	return ID(id)
}

type IDs []ID

func NewIDs(
	id ...ID,
) IDs {
	ids := make(IDs, 0)
	ids = append(ids, id...)
	return ids
}
