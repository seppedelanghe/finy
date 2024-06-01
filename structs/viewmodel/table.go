package viewmodel

type TableVM struct {
	Headers []TableHeaderVM
	Rows []TableRow
}

type TableHeaderVM struct {
	Name string
}

type TableRow struct {
	Data []TableDataVM
}

type TableDataVM struct {
	Value string
	Bold bool
}

