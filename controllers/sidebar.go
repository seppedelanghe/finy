package controllers

import (
	"net/http"
	"strconv"

	"finy.be/api/rendering"
	"finy.be/api/structs/viewmodel"
)

func GetSidebar(w http.ResponseWriter, r *http.Request) {
	sidebar := viewmodel.SidebarVM {
		Selected: 1,
		SelectedMenu: 0,
		Pages: []viewmodel.SidebarPageVM {
			{ Icon: "bell.svg", Name: "Notifications" },
			{ Icon: "chart.svg", Name: "Accounts" },
			{ Icon: "settings.svg", Name: "Settings" },
		},
		Menus: []viewmodel.SidebarMenuVM {
			{ Title: "Checking", Subtitle: "1114.13", Icon: "" },
			{ Title: "Savings", Subtitle: "9876.10", Icon: "" },
		},
	}

	rendering.Renderer.HTML(w, http.StatusOK, "sidebar", sidebar)
}

func PostSidebar(w http.ResponseWriter, r *http.Request) {
	selected := r.PathValue("index")
	index, err := strconv.Atoi(selected)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	sidebar := viewmodel.SidebarVM {
		Selected: 1,
		SelectedMenu: index,
		Pages: []viewmodel.SidebarPageVM {
			{ Icon: "bell.svg", Name: "Notifications" },
			{ Icon: "chart.svg", Name: "Accounts" },
			{ Icon: "settings.svg", Name: "Settings" },
		},
		Menus: []viewmodel.SidebarMenuVM {
			{ Title: "Checking", Subtitle: "1114.13", Icon: "" },
			{ Title: "Savings", Subtitle: "9876.10", Icon: "" },
		},
	}

	rendering.Renderer.HTML(w, http.StatusOK, "sidebar", sidebar)
}

