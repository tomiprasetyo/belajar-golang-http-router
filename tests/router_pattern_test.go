package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			id := p.ByName("id")
			itemId := p.ByName("itemId")
			text := "Product " + id + " Item " + itemId
			fmt.Fprint(w, text)
		})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/products/5/items/10", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Product 5 Item 10", string(body))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image",
		func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			image := p.ByName("image")
			text := "Image : " + image
			fmt.Fprint(w, text)
		})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/images/large/book.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image : /large/book.png", string(body))
}
