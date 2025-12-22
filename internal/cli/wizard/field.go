package wizard

type Field struct {
	Label    string
	Value    func() string
	SetValue func(string) error
	Optional bool
}
