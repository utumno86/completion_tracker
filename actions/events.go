package actions

import (
    "net/http"

    "github.com/gobuffalo/buffalo"
)

type EventsResource struct{
    buffalo.Resource
}


// List default implementation.
func (v EventsResource) List(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Event#List"))
}

// Show default implementation.
func (v EventsResource) Show(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Event#Show"))
}

// Create default implementation.
func (v EventsResource) Create(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Event#Create"))
}

// Update default implementation.
func (v EventsResource) Update(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Event#Update"))
}

// Destroy default implementation.
func (v EventsResource) Destroy(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Event#Destroy"))
}

// New default implementation.
func (v EventsResource) New(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Event#New"))
}

// Edit default implementation.
func (v EventsResource) Edit(c buffalo.Context) error {
    return c.Render(http.StatusOK, r.String("Event#Edit"))
}

