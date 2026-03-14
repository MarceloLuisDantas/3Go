package logic

type Group struct {
	Cells [][]int
}

type Groups struct {
	Groups []Group
}

func NewGroups() *Groups {
	return &Groups{}
}
