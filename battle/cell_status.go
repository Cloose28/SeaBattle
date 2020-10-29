package battle

type cellStatus int

const (
	Empty cellStatus = 0
	Shoted cellStatus = -1
)

func (c cellStatus) Int() int {
	return int(c)
}


