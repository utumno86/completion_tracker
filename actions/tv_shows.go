package actions

import (
    "fmt"
    "net/http"

    "completion_tracker/models"

    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/pop/v6"
    "github.com/gobuffalo/x/responder"
)

type TvShowsResource struct{
    buffalo.Resource
}


// List gets all TV Show completions. This function is mapped to the path
// GET /tv_shows
func (v TvShowsResource) List(c buffalo.Context) error {
    // Get the DB connection from the context
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completions := &models.Completions{}

    // Paginate results and filter by TV Show type
    q := tx.PaginateFromParams(c.Params())
    q = q.Where("type = ?", models.CompletionTypeTVShow)

    // Retrieve all TV Show Completions from the DB
    if err := q.All(completions); err != nil {
        return err
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        c.Set("pagination", q.Paginator)
        c.Set("completions", completions)
        return c.Render(http.StatusOK, r.HTML("tv_shows/index.plush.html"))
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(200, r.JSON(completions))
    }).Wants("xml", func(c buffalo.Context) error {
        return c.Render(200, r.XML(completions))
    }).Respond(c)
}

// Show gets the data for one TV Show completion. This function is mapped to
// the path GET /tv_shows/{tv_show_id}
func (v TvShowsResource) Show(c buffalo.Context) error {
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completion := &models.Completion{}
    if err := tx.Find(completion, c.Param("tv_show_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    // Ensure this is actually a TV Show completion
    if completion.Type != models.CompletionTypeTVShow {
        return c.Error(http.StatusNotFound, fmt.Errorf("completion is not a TV show"))
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        c.Set("completion", completion)
        return c.Render(http.StatusOK, r.HTML("tv_shows/show.plush.html"))
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(200, r.JSON(completion))
    }).Wants("xml", func(c buffalo.Context) error {
        return c.Render(200, r.XML(completion))
    }).Respond(c)
}

// Create adds a TV Show completion to the DB. This function is mapped to the
// path POST /tv_shows
func (v TvShowsResource) Create(c buffalo.Context) error {
    completion := &models.Completion{
        Type: models.CompletionTypeTVShow,
    }

    if err := c.Bind(completion); err != nil {
        return err
    }

    // Ensure type is set correctly
    completion.Type = models.CompletionTypeTVShow

    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    verrs, err := tx.ValidateAndCreate(completion)
    if err != nil {
        return err
    }

    if verrs.HasAny() {
        return responder.Wants("html", func(c buffalo.Context) error {
            c.Set("errors", verrs)
            c.Set("completion", completion)
            return c.Render(http.StatusUnprocessableEntity, r.HTML("tv_shows/new.plush.html"))
        }).Wants("json", func(c buffalo.Context) error {
            return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
        }).Wants("xml", func(c buffalo.Context) error {
            return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
        }).Respond(c)
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        c.Flash().Add("success", T.Translate(c, "completion.created.success"))
        return c.Redirect(http.StatusSeeOther, "/tv_shows/%v", completion.ID)
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(http.StatusCreated, r.JSON(completion))
    }).Wants("xml", func(c buffalo.Context) error {
        return c.Render(http.StatusCreated, r.XML(completion))
    }).Respond(c)
}

// Update changes a TV Show completion in the DB. This function is mapped to
// the path PUT /tv_shows/{tv_show_id}
func (v TvShowsResource) Update(c buffalo.Context) error {
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completion := &models.Completion{}
    if err := tx.Find(completion, c.Param("tv_show_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    if completion.Type != models.CompletionTypeTVShow {
        return c.Error(http.StatusNotFound, fmt.Errorf("completion is not a TV show"))
    }

    if err := c.Bind(completion); err != nil {
        return err
    }

    // Ensure type remains TV Show
    completion.Type = models.CompletionTypeTVShow

    verrs, err := tx.ValidateAndUpdate(completion)
    if err != nil {
        return err
    }

    if verrs.HasAny() {
        return responder.Wants("html", func(c buffalo.Context) error {
            c.Set("errors", verrs)
            c.Set("completion", completion)
            return c.Render(http.StatusUnprocessableEntity, r.HTML("tv_shows/edit.plush.html"))
        }).Wants("json", func(c buffalo.Context) error {
            return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
        }).Wants("xml", func(c buffalo.Context) error {
            return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
        }).Respond(c)
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        c.Flash().Add("success", T.Translate(c, "completion.updated.success"))
        return c.Redirect(http.StatusSeeOther, "/tv_shows/%v", completion.ID)
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(http.StatusOK, r.JSON(completion))
    }).Wants("xml", func(c buffalo.Context) error {
        return c.Render(http.StatusOK, r.XML(completion))
    }).Respond(c)
}

// Destroy deletes a TV Show completion from the DB. This function is mapped
// to the path DELETE /tv_shows/{tv_show_id}
func (v TvShowsResource) Destroy(c buffalo.Context) error {
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completion := &models.Completion{}
    if err := tx.Find(completion, c.Param("tv_show_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    if completion.Type != models.CompletionTypeTVShow {
        return c.Error(http.StatusNotFound, fmt.Errorf("completion is not a TV show"))
    }

    if err := tx.Destroy(completion); err != nil {
        return err
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        c.Flash().Add("success", T.Translate(c, "completion.destroyed.success"))
        return c.Redirect(http.StatusSeeOther, "/tv_shows")
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(http.StatusOK, r.JSON(completion))
    }).Wants("xml", func(c buffalo.Context) error {
        return c.Render(http.StatusOK, r.XML(completion))
    }).Respond(c)
}

// New renders the form for creating a new TV Show completion.
// This function is mapped to the path GET /tv_shows/new
func (v TvShowsResource) New(c buffalo.Context) error {
    completion := &models.Completion{
        Type: models.CompletionTypeTVShow,
    }
    c.Set("completion", completion)

    return c.Render(http.StatusOK, r.HTML("tv_shows/new.plush.html"))
}

// Edit renders a edit form for a TV Show completion. This function is
// mapped to the path GET /tv_shows/{tv_show_id}/edit
func (v TvShowsResource) Edit(c buffalo.Context) error {
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completion := &models.Completion{}
    if err := tx.Find(completion, c.Param("tv_show_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    if completion.Type != models.CompletionTypeTVShow {
        return c.Error(http.StatusNotFound, fmt.Errorf("completion is not a TV show"))
    }

    c.Set("completion", completion)
    return c.Render(http.StatusOK, r.HTML("tv_shows/edit.plush.html"))
}

