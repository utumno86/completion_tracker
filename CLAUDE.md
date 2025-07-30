# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Buffalo web application called "completion_tracker" built with Go for tracking completion metrics and progress over time. Buffalo is a Go web framework that provides rapid web development with hot reloading, asset pipeline, and built-in ORM (Pop).

### Core Features
- **Completion Management**: Full CRUD operations for completion records
- **Data Model**: Tracks name, completion count, and completion timestamp
- **Responsive UI**: Bootstrap 5 interface with navigation
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
  - `completions.go` - Completion CRUD handlers
- `models/` - Database models and Pop ORM configuration
  - `completion.go` - Completion model with validation
- `templates/` - Plush HTML templates
  - `completions/` - Completion views (index, show, new, edit, _form)
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
   - `/completions` - Completion resource routes (RESTful CRUD)
5. Database connection is initialized in `models/models.go`

### Data Models
- **Completion**: Tracks completion records with fields:
  - `Name` (string) - Required
  - `Completions` (int) - Required, number of completions
  - `CompletedAt` (timestamp) - Required, when completed
  - Standard UUID ID, CreatedAt, UpdatedAt fields

### Routes
The application uses Buffalo's resource routing for completions:
- `GET /completions` - List all completions (with pagination)
- `GET /completions/new` - Show creation form
- `POST /completions` - Create new completion
- `GET /completions/{id}` - Show specific completion
- `GET /completions/{id}/edit` - Show edit form
- `PUT /completions/{id}` - Update completion
- `DELETE /completions/{id}` - Delete completion

All routes support HTML, JSON, and XML content types.

### Database Configuration
Database settings are in `database.yml` with separate configs for development, test, and production environments. Uses PostgreSQL by default.

### Frontend Assets
Assets are managed by Webpack with live reloading. SCSS and JS files in `assets/` are compiled to `public/assets/`. Navigation bar includes links to Home and Completions.