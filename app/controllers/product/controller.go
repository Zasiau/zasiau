package product

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mholt/binding"

	"github.com/dongri/gonion/app/middlewares/render"
	"github.com/dongri/gonion/app/services"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	const view = "product/index.html"
	products, err := services.GetProducts(r)
	if err != nil {
		log.Println(err)
		return
	}
	output := map[string]interface{}{
		"products": products,
	}
	render.HTML(w, r, view, output)
}

// Show ...
func Show(w http.ResponseWriter, r *http.Request) {
	const view = "product/show.html"
	ID := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	product, err := services.GetProductByID(r, id)
	if err != nil {
		log.Println(err)
		return
	}
	output := map[string]interface{}{
		"product": product,
	}
	render.HTML(w, r, view, output)
}

// New ...
func New(w http.ResponseWriter, r *http.Request) {
	const view = "product/new.html"
	output := map[string]interface{}{}
	render.HTML(w, r, view, output)
}

// Edit ...
func Edit(w http.ResponseWriter, r *http.Request) {
	const view = "product/edit.html"
	ID := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	product, err := services.GetProductByID(r, id)
	if err != nil {
		log.Println(err)
		return
	}
	output := map[string]interface{}{
		"product": product,
	}
	render.HTML(w, r, view, output)
}

// ProductForm ...
type ProductForm struct {
	Name  string
	Price string
}

// FieldMap ...
func (s *ProductForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&s.Name: binding.Field{
			Form:     "name",
			Required: true,
		},
		&s.Price: binding.Field{
			Form:     "price",
			Required: true,
		},
	}
}

// Create ...
func Create(w http.ResponseWriter, r *http.Request) {
	form := new(ProductForm)
	if err := binding.Bind(r, form); err != nil {
		log.Println(err)
		return
	}
	name := form.Name
	price := form.Price
	uint32Price, err := strconv.ParseUint(price, 10, 32)
	if err != nil {
		log.Println(err)
		return
	}
	if err := services.CreateProduct(r, name, uint32(uint32Price)); err != nil {
		log.Println(err)
		return
	}
	http.Redirect(w, r, "/products", http.StatusFound)
}

// Update ...
func Update(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		log.Println(err)
		return
	}
	form := new(ProductForm)
	if err := binding.Bind(r, form); err != nil {
		log.Println(err)
		return
	}
	name := form.Name
	price := form.Price
	uint32Price, err := strconv.ParseUint(price, 10, 32)
	if err != nil {
		log.Println(err)
		return
	}
	if err := services.UpdateProductByID(r, id, name, uint32(uint32Price)); err != nil {
		log.Println(err)
		return
	}
	http.Redirect(w, r, "/products", http.StatusFound)
}

// Destroy ...
func Destroy(w http.ResponseWriter, r *http.Request) {

}
