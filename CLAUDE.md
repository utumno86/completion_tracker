# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Buffalo web application called "completion_tracker" built with Go for tracking completion metrics and progress over time. Buffalo is a Go web framework that provides rapid web development with hot reloading, asset pipeline, and built-in ORM (Pop).

### Core Features
- **Type-Specific Completion Management**: Specialized interfaces for different completion types
  - TV Shows (episodes watched), Video Games (hours played), Books (reading progress)
  - Audio Books (listening hours), Events (participation tracking)
- **Automatic Type Detection**: Interface determines completion type based on route
- **Unified Data Model**: Single completion table with type differentiation
- **Responsive UI**: Bootstrap 5 interface with dropdown navigation
- **API Support**: JSON/XML endpoints alongside HTML views

## Development Commands

### Starting the Application
```bash
buffalo dev
```
Starts the development server with hot reloading on port 3000.

### Database Operations
```bash
buffalo pop create -a       # Create all databases
buffalo pop migrate         # Run migrations
buffalo pop reset           # Reset database
```

### Asset Building
```bash
npm run dev                 # Watch and rebuild assets during development
npm run build              # Build assets for production
```

### Testing
```bash
go test ./...              # Run all Go tests
go test ./actions          # Test specific package
```

### Building for Production
```bash
buffalo build              # Build the complete application binary
```

## Architecture

### Directory Structure
- `cmd/app/main.go` - Application entry point
- `actions/` - HTTP handlers and routes (Buffalo's controllers)
  - `app.go` - Main application configuration and middleware setup
  - `home.go` - Home page handler
  - `completions.go` - General completion CRUD handlers
  - `tv_shows.go` - TV show-specific completion handlers
  - `video_games.go` - Video game-specific completion handlers
  - `books.go`, `audio_books.go`, `events.go` - Other type-specific handlers
- `models/` - Database models and Pop ORM configuration
  - `completion.go` - Unified completion model with type system and validation
- `templates/` - Plush HTML templates
  - `completions/` - General completion views
  - `tv_shows/`, `video_games/`, `books/`, `audio_books/`, `events/` - Type-specific views
- `assets/` - Frontend assets (SCSS, JS) processed by Webpack
- `public/` - Static files and compiled assets
- `locales/` - Internationalization files
- `migrations/` - Database migration files
- `grifts/` - Task runners (Buffalo's equivalent of Rake tasks)

### Key Technologies
- **Buffalo Framework** - Go web framework
- **Pop ORM** - Database ORM for Go
- **Plush Templates** - Go templating engine
- **PostgreSQL** - Primary database
- **Webpack** - Asset bundling
- **Bootstrap 5** - CSS framework

### Application Flow
1. `cmd/app/main.go` starts the application
2. `actions/app.go` configures middleware and routes
3. Middleware stack includes CSRF protection, SSL forcing, database transactions
4. Routes are defined in `App()` function in `actions/app.go`:
   - `/` - Home page
   - `/completions` - General completion resource routes (RESTful CRUD)
   - `/tv_shows`, `/video_games`, `/books`, `/audio_books`, `/events` - Type-specific routes
5. Database connection is initialized in `models/models.go`

### Data Models
- **Completion**: Unified model for all completion types with fields:
  - `Name` (string) - Required, name of the item being completed
  - `Type` (CompletionType) - Required, one of: TV Show, Video Game, Book, Audio Book, Event
  - `Completions` (int) - Required, count value (episodes, hours, pages, etc.)
  - `CompletedAt` (timestamp) - Required, when completed
  - Standard UUID ID, CreatedAt, UpdatedAt fields
- **CompletionType**: Enum with 5 predefined types and validation

### Routes
The application uses Buffalo's resource routing with both general and type-specific endpoints:

**General Routes**:
- `GET /completions` - List all completions (all types, with pagination)
- `GET /completions/new` - Show general creation form
- `POST /completions` - Create new completion
- `GET /completions/{id}` - Show specific completion
- `GET /completions/{id}/edit` - Show edit form
- `PUT /completions/{id}` - Update completion
- `DELETE /completions/{id}` - Delete completion

**Type-Specific Routes**:
- `GET /tv_shows` - List TV show completions only
- `GET /tv_shows/new` - Show TV show creation form (type auto-set)
- `POST /tv_shows` - Create TV show completion
- `GET /tv_shows/{id}` - Show specific TV show
- `GET /tv_shows/{id}/edit` - Edit TV show completion
- `PUT /tv_shows/{id}` - Update TV show completion
- `DELETE /tv_shows/{id}` - Delete TV show completion

*(Similar patterns exist for `/video_games`, `/books`, `/audio_books`, `/events`)*

All routes support HTML, JSON, and XML content types.

### Database Configuration
Database settings are in `database.yml` with separate configs for development, test, and production environments. Uses PostgreSQL by default.

### Frontend Assets
Assets are managed by Webpack with live reloading. SCSS and JS files in `assets/` are compiled to `public/assets/`. 

### User Interface
- **Navigation**: Bootstrap dropdown menu with links to each completion type
- **Type-Specific Forms**: Each completion type has specialized input forms with relevant terminology
- **Progress Display**: Type-appropriate status indicators and progress visualization
- **Template Conventions**: Buffalo uses `_form.html` naming for form partials referenced as `partial("resource/form.html")`

### Type System Architecture
- **Single Table**: All completions stored in one `completions` table with `type` column
- **Type Safety**: Controllers validate and enforce correct types for each endpoint
- **Automatic Type Setting**: Type-specific controllers automatically set the appropriate type
- **Specialized Interfaces**: Each type has tailored UI with relevant field names and validation