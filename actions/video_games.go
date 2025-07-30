package actions

import (
    "fmt"
    "net/http"

    "completion_tracker/models"

    "github.com/gobuffalo/buffalo"
    "github.com/gobuffalo/pop/v6"
    "github.com/gobuffalo/x/responder"
)

type VideoGamesResource struct{
    buffalo.Resource
}

// List gets all Video Game completions
func (v VideoGamesResource) List(c buffalo.Context) error {
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completions := &models.Completions{}
    q := tx.PaginateFromParams(c.Params())
    q = q.Where("type = ?", models.CompletionTypeVideoGame)

    if err := q.All(completions); err != nil {
        return err
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        c.Set("pagination", q.Paginator)
        c.Set("completions", completions)
        return c.Render(http.StatusOK, r.HTML("video_games/index.plush.html"))
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(200, r.JSON(completions))
    }).Respond(c)
}

// Show gets the data for one Video Game completion
func (v VideoGamesResource) Show(c buffalo.Context) error {
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completion := &models.Completion{}
    if err := tx.Find(completion, c.Param("video_game_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    if completion.Type != models.CompletionTypeVideoGame {
        return c.Error(http.StatusNotFound, fmt.Errorf("completion is not a video game"))
    }

    return responder.Wants("html", func(c buffalo.Context) error {
        c.Set("completion", completion)
        return c.Render(http.StatusOK, r.HTML("video_games/show.plush.html"))
    }).Wants("json", func(c buffalo.Context) error {
        return c.Render(200, r.JSON(completion))
    }).Respond(c)
}

// New renders the form for creating a new Video Game completion
func (v VideoGamesResource) New(c buffalo.Context) error {
    completion := &models.Completion{
        Type: models.CompletionTypeVideoGame,
    }
    c.Set("completion", completion)
    return c.Render(http.StatusOK, r.HTML("video_games/new.plush.html"))
}

// Create adds a Video Game completion to the DB
func (v VideoGamesResource) Create(c buffalo.Context) error {
    completion := &models.Completion{
        Type: models.CompletionTypeVideoGame,
    }

    if err := c.Bind(completion); err != nil {
        return err
    }
    completion.Type = models.CompletionTypeVideoGame

    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    verrs, err := tx.ValidateAndCreate(completion)
    if err != nil {
        return err
    }

    if verrs.HasAny() {
        c.Set("errors", verrs)
        c.Set("completion", completion)
        return c.Render(http.StatusUnprocessableEntity, r.HTML("video_games/new.plush.html"))
    }

    c.Flash().Add("success", T.Translate(c, "completion.created.success"))
    return c.Redirect(http.StatusSeeOther, "/video_games/%v", completion.ID)
}

// Edit renders a edit form for a Video Game completion
func (v VideoGamesResource) Edit(c buffalo.Context) error {
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completion := &models.Completion{}
    if err := tx.Find(completion, c.Param("video_game_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    if completion.Type != models.CompletionTypeVideoGame {
        return c.Error(http.StatusNotFound, fmt.Errorf("completion is not a video game"))
    }

    c.Set("completion", completion)
    return c.Render(http.StatusOK, r.HTML("video_games/edit.plush.html"))
}

// Update changes a Video Game completion in the DB
func (v VideoGamesResource) Update(c buffalo.Context) error {
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completion := &models.Completion{}
    if err := tx.Find(completion, c.Param("video_game_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    if completion.Type != models.CompletionTypeVideoGame {
        return c.Error(http.StatusNotFound, fmt.Errorf("completion is not a video game"))
    }

    if err := c.Bind(completion); err != nil {
        return err
    }
    completion.Type = models.CompletionTypeVideoGame

    verrs, err := tx.ValidateAndUpdate(completion)
    if err != nil {
        return err
    }

    if verrs.HasAny() {
        c.Set("errors", verrs)
        c.Set("completion", completion)
        return c.Render(http.StatusUnprocessableEntity, r.HTML("video_games/edit.plush.html"))
    }

    c.Flash().Add("success", T.Translate(c, "completion.updated.success"))
    return c.Redirect(http.StatusSeeOther, "/video_games/%v", completion.ID)
}

// Destroy deletes a Video Game completion from the DB
func (v VideoGamesResource) Destroy(c buffalo.Context) error {
    tx, ok := c.Value("tx").(*pop.Connection)
    if !ok {
        return fmt.Errorf("no transaction found")
    }

    completion := &models.Completion{}
    if err := tx.Find(completion, c.Param("video_game_id")); err != nil {
        return c.Error(http.StatusNotFound, err)
    }

    if completion.Type != models.CompletionTypeVideoGame {
        return c.Error(http.StatusNotFound, fmt.Errorf("completion is not a video game"))
    }

    if err := tx.Destroy(completion); err != nil {
        return err
    }

    c.Flash().Add("success", T.Translate(c, "completion.destroyed.success"))
    return c.Redirect(http.StatusSeeOther, "/video_games")
}

