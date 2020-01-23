package builder

type Builder interface {
	MakeTitle(title string) error
	MakeString(str string)
	MakeItems(items []string)
	Close() error
}
