package actions

import (
    "fmt"
    "net/http"

    "completion_tracker/models"

    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/pop/v6"
    "github.com/gobuffalo/x/responder"
)

type CompletionsResource struct{
    buffalo.Resource
}

// List gets all Completions. This function is mapped to the path
// GET /completions
func (v CompletionsResource) List(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completions := &models.Completions{}

    // Paginate results. Params "page" and "per_page" control pagination.
    // Default values are "page=1" and "per_page=20".
    q := tx.PaginateFromParams(c.Params())

    // Retrieve all Completions from the DB
    if err := q.All(completions); err != nil {
        return err
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        // Add the paginator to the context so it can be used in the template.
        c.Set("pagination", q.Paginator)

        c.Set("completions", completions)
        return c.Render(http.StatusOK, r.HTML("completions/index.plush.html"))
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(200, r.JSON(completions))
    }).Wants("xml", func(c buffalo.Context) error {
        return c.Render(200, r.XML(completions))
    }).Respond(c)
}

// Show gets the data for one Completion. This function is mapped to
// the path GET /completions/{completion_id}
func (v CompletionsResource) Show(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    // Allocate an empty Completion
    completion := &models.Completion{}

    // To find the Completion the parameter completion_id is used.
    if err := tx.Find(completion, c.Param("completion_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        c.Set("completion", completion)

        return c.Render(http.StatusOK, r.HTML("completions/show.plush.html"))
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(200, r.JSON(completion))
    }).Wants("xml", func(c buffalo.Context) error {
        return c.Render(200, r.XML(completion))
    }).Respond(c)
}

// New renders the form for creating a new Completion.
// This function is mapped to the path GET /completions/new
func (v CompletionsResource) New(c buffalo.Context) error {
    c.Set("completion", &models.Completion{})

    return c.Render(http.StatusOK, r.HTML("completions/new.plush.html"))
}

// Create adds a Completion to the DB. This function is mapped to the
// path POST /completions
func (v CompletionsResource) Create(c buffalo.Context) error {
    // Allocate an empty Completion
    completion := &models.Completion{}

    // Bind completion to the html form elements
    if err := c.Bind(completion); err != nil {
        return err
    }

    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    // Validate the data from the html form
    verrs, err := tx.ValidateAndCreate(completion)
    if err != nil {
        return err
    }

    if verrs.HasAny() {
        return responder.Wants("html", func(c buffalo.Context) error {
            // Make the errors available inside the html template
            c.Set("errors", verrs)

            // Render again the new.html template that the user can
            // correct the input.
            c.Set("completion", completion)

            return c.Render(http.StatusUnprocessableEntity, r.HTML("completions/new.plush.html"))
        }).Wants("json", func(c buffalo.Context) error {
            return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
        }).Wants("xml", func(c buffalo.Context) error {
            return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
        }).Respond(c)
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        // If there are no errors set a success message
        c.Flash().Add("success", T.Translate(c, "completion.created.success"))

        // and redirect to the show page
        return c.Redirect(http.StatusSeeOther, "/completions/%v", completion.ID)
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(http.StatusCreated, r.JSON(completion))
    }).Wants("xml", func(c buffalo.Context) error {
        return c.Render(http.StatusCreated, r.XML(completion))
    }).Respond(c)
}

// Edit renders a edit form for a Completion. This function is
// mapped to the path GET /completions/{completion_id}/edit
func (v CompletionsResource) Edit(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    // Allocate an empty Completion
    completion := &models.Completion{}

    if err := tx.Find(completion, c.Param("completion_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    c.Set("completion", completion)
    return c.Render(http.StatusOK, r.HTML("completions/edit.plush.html"))
}

// Update changes a Completion in the DB. This function is mapped to
// the path PUT /completions/{completion_id}
func (v CompletionsResource) Update(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    // Allocate an empty Completion
    completion := &models.Completion{}

    if err := tx.Find(completion, c.Param("completion_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    // Bind Completion to the html form elements
    if err := c.Bind(completion); err != nil {
        return err
    }

    verrs, err := tx.ValidateAndUpdate(completion)
    if err != nil {
        return err
    }

    if verrs.HasAny() {
        return responder.Wants("html", func(c buffalo.Context) error {
            // Make the errors available inside the html template
            c.Set("errors", verrs)

            // Render again the edit.html template that the user can
            // correct the input.
            c.Set("completion", completion)

            return c.Render(http.StatusUnprocessableEntity, r.HTML("completions/edit.plush.html"))
        }).Wants("json", func(c buffalo.Context) error {
            return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
        }).Wants("xml", func(c buffalo.Context) error {
            return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
        }).Respond(c)
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        // If there are no errors set a success message
        c.Flash().Add("success", T.Translate(c, "completion.updated.success"))

        // and redirect to the show page
        return c.Redirect(http.StatusSeeOther, "/completions/%v", completion.ID)
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(http.StatusOK, r.JSON(completion))
    }).Wants("xml", func(c buffalo.Context) error {
        return c.Render(http.StatusOK, r.XML(completion))
    }).Respond(c)
}

// Destroy deletes a Completion from the DB. This function is mapped
// to the path DELETE /completions/{completion_id}
func (v CompletionsResource) Destroy(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    // Allocate an empty Completion
    completion := &models.Completion{}

    // To find the Completion the parameter completion_id is used.
    if err := tx.Find(completion, c.Param("completion_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    if err := tx.Destroy(completion); err != nil {
        return err
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        // If there are no errors set a flash message
        c.Flash().Add("success", T.Translate(c, "completion.destroyed.success"))

        // Redirect to the index page
        return c.Redirect(http.StatusSeeOther, "/completions")
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(http.StatusOK, r.JSON(completion))
    }).Wants("xml", func(c buffalo.Context) error {
        return c.Render(http.StatusOK, r.XML(completion))
    }).Respond(c)
}

