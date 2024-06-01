package viewmodel

type TableVM struct {
	Headers []TableHeaderVM
	Items []TableItemVM
}

type TableHeaderVM struct {
	Name string
}

type TableItemVM struct {
	Value string
	Bold bool
}

