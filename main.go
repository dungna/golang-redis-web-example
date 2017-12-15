package main

import (
	"net/http"
	"strconv"
	"./models"
	"fmt"
)

func main() {
	http.HandleFunc("/album", showAlbum)
	http.HandleFunc("/like", addLike)
	http.HandleFunc("/popular", listPopular)

	http.ListenAndServe(":3000", nil)
}

func showAlbum(w http.ResponseWriter, r *http.Request)  {
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, http.StatusText(405), 405)
		return
	}
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	bk, err := models.FindAlbum(id)
	if err == models.ErrNoAlbum {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	// Write the album details as plain text to the client.
	fmt.Fprintf(w, "%s by %s: £%.2f [%d likes] \n", bk.Title, bk.Artist, bk.Price, bk.Likes)
}

func addLike(w http.ResponseWriter, r *http.Request) {
	// Unless the request is using the POST method, return a 405 'Method Not
	// Allowed' response.
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, http.StatusText(405), 405)
		return
	}

	// Retreive the id from the POST request body. If there is no parameter
	// named "id" in the request body then PostFormValue() will return an
	// empty string. We check for this, returning a 400 Bad Request response
	// if it's missing.
	id := r.PostFormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	// Validate that the id is a valid integer by trying to convert it,
	// returning a 400 Bad Request response if the conversion fails.
	if _, err := strconv.Atoi(id); err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	// Call the IncrementLikes() function passing in the user-provided id. If
	// there's no album found with that id, return a 404 Not Found response.
	// In the event of any other errors, return a 500 Internal Server Error
	// response.
	err := models.IncrementLikes(id)
	if err == models.ErrNoAlbum {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// Redirect the client to the GET /ablum route, so they can see the
	// impact their like has had.
	http.Redirect(w, r, "/album?id="+id, 303)
}

func listPopular(w http.ResponseWriter, r *http.Request) {
	// Unless the request is using the GET method, return a 405 'Method Not
	// Allowed' response.
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		http.Error(w, http.StatusText(405), 405)
		return
	}

	// Call the FindTopThree() function, returning a return a 500 Internal
	// Server Error response if there's any error.
	abs, err := models.FindTopThree()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	// Loop through the 3 albums, writing the details as a plain text list
	// to the client.
	for i, ab := range abs {
		fmt.Fprintf(w, "%d) %s by %s: £%.2f [%d likes] \n", i+1, ab.Title, ab.Artist, ab.Price, ab.Likes)
	}
}