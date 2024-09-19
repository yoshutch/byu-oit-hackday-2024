package adapters

import (
	"byu.edu/hackday-favorite-color/clients"
	dto2 "byu.edu/hackday-favorite-color/dto"
	"byu.edu/hackday-favorite-color/services"
	"byu.edu/hackday-profile/dto"
	"html/template"
	"log"
	"net/http"
)

type HtmlAdapter struct {
	mux             *http.ServeMux
	indexTmpl       *template.Template
	favColorService *services.FavColorService
}

func NewHtmlAdapter(mux *http.ServeMux, service *services.FavColorService) (*HtmlAdapter, error) {
	// TODO parse all templates into a map of some kind?
	indexTmpl, err := template.ParseFiles("pages/index.html", "pages/layout.html")
	if err != nil {
		log.Printf("Error parsing file: %s", err)
		return nil, err
	}
	if err != nil {
		log.Printf("Error parsing file: %s", err)
		return nil, err
	}
	return &HtmlAdapter{
		mux,
		indexTmpl,
		service,
	}, nil
}

func (h HtmlAdapter) HandleRoutes() {
	// Index.html
	h.mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		type IndexData struct {
			Profile  *dto.Profile
			FavColor *dto2.FavoriteColor
		}

		profile, err := clients.GetProfile(1)
		if err != nil {
			log.Printf("Error calling API: %s", err)
			return
		}
		favColor, err := h.favColorService.LoadFavColor(1)
		if err != nil {
			return
		}
		data := IndexData{
			Profile:  profile,
			FavColor: favColor,
		}
		err = h.indexTmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Printf("Error executing template: %s", err)
			return
		}
	})
}
