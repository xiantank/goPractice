package KeyCount

type KeyCount interface {
	Add(string)
	Total() int
	Print()
	PrintType() string
}
type KeyNode struct {
	Key   string
	Count int
}
