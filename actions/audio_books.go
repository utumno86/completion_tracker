package actions

import (
    "net/http"

    "github.com/gobuffalo/buffalo"
)

type AudioBooksResource struct{
    buffalo.Resource
}


// List default implementation.
func (v AudioBooksResource) List(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("AudioBook#List"))
}

// Show default implementation.
func (v AudioBooksResource) Show(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("AudioBook#Show"))
}

// Create default implementation.
func (v AudioBooksResource) Create(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("AudioBook#Create"))
}

// Update default implementation.
func (v AudioBooksResource) Update(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("AudioBook#Update"))
}

// Destroy default implementation.
func (v AudioBooksResource) Destroy(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("AudioBook#Destroy"))
}

// New default implementation.
func (v AudioBooksResource) New(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("AudioBook#New"))
}

// Edit default implementation.
func (v AudioBooksResource) Edit(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("AudioBook#Edit"))
}

