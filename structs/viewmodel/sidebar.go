package viewmodel

type SidebarVM struct {
	Title string
	Selected int
	Pages []SidebarPageVM
	Menus []SidebarMenuVM
}

type SidebarPageVM struct {
	Icon string
	Name string
}


type SidebarMenuVM struct {
	Icon string
	Title string
	Subtitle string
}

