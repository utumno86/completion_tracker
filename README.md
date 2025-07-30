# Completion Tracker

A Buffalo web application for tracking completion metrics and progress over time.

## Features

- **Type-Specific Completion Tracking**: Specialized interfaces for different completion types
  - 📺 **TV Shows**: Track episodes watched with show-specific terminology
  - 🎮 **Video Games**: Track hours played with gaming-focused interface  
  - 📚 **Books**: Track reading progress with book-specific fields
  - 🎧 **Audio Books**: Track listening hours with audio-specific interface
  - 📅 **Events**: Track event attendance and participation
- **Automatic Type Detection**: Interface determines completion type automatically
- **Full CRUD Operations**: Create, read, update, and delete completion entries
- **Responsive UI**: Bootstrap 5 based interface with dropdown navigation
- **Data Validation**: Form validation with error messaging
- **API Support**: JSON and XML endpoints alongside HTML views

## Database Setup

This application uses PostgreSQL. Make sure PostgreSQL is running and update the `database.yml` file with your database credentials.

### Create Your Databases

```console
buffalo pop create -a
buffalo pop migrate
```

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

```console
buffalo dev
```

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see the Completion Tracker homepage.

## Usage

### Web Interface
- **Home**: [http://127.0.0.1:3000](http://127.0.0.1:3000)
- **Type-Specific Interfaces**:
  - **TV Shows**: [http://127.0.0.1:3000/tv_shows](http://127.0.0.1:3000/tv_shows) - Track episodes watched
  - **Video Games**: [http://127.0.0.1:3000/video_games](http://127.0.0.1:3000/video_games) - Track hours played
  - **Books**: [http://127.0.0.1:3000/books](http://127.0.0.1:3000/books) - Track reading progress
  - **Audio Books**: [http://127.0.0.1:3000/audio_books](http://127.0.0.1:3000/audio_books) - Track listening hours
  - **Events**: [http://127.0.0.1:3000/events](http://127.0.0.1:3000/events) - Track event participation
- **All Completions**: [http://127.0.0.1:3000/completions](http://127.0.0.1:3000/completions) - Unified view of all types

Each interface provides:
- Specialized forms with relevant terminology
- Type-appropriate progress displays
- Contextual status indicators

### API Endpoints
All endpoints support JSON and XML content types:

**General Completions**:
- `GET /completions` - List all completions (all types)
- `GET /completions/{id}` - Get specific completion
- `POST /completions` - Create new completion
- `PUT /completions/{id}` - Update completion
- `DELETE /completions/{id}` - Delete completion

**Type-Specific Endpoints**:
- `GET /tv_shows` - List TV show completions
- `POST /tv_shows` - Create TV show completion
- `GET /tv_shows/{id}` - Get specific TV show
- `PUT /tv_shows/{id}` - Update TV show
- `DELETE /tv_shows/{id}` - Delete TV show

*(Similar patterns available for `/video_games`, `/books`, `/audio_books`, `/events`)*

## Development

### Running Tests
```console
go test ./...
```

### Building for Production
```console
buffalo build
```

[Powered by Buffalo](http://gobuffalo.io)
