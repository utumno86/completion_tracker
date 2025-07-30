package actions

import (
    "net/http"

    "github.com/gobuffalo/buffalo"
)

type BooksResource struct{
    buffalo.Resource
}


// List default implementation.
func (v BooksResource) List(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Book#List"))
}

// Show default implementation.
func (v BooksResource) Show(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Book#Show"))
}

// Create default implementation.
func (v BooksResource) Create(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Book#Create"))
}

// Update default implementation.
func (v BooksResource) Update(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Book#Update"))
}

// Destroy default implementation.
func (v BooksResource) Destroy(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Book#Destroy"))
}

// New default implementation.
func (v BooksResource) New(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Book#New"))
}

// Edit default implementation.
func (v BooksResource) Edit(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Book#Edit"))
}

