# Completion Tracker

A Buffalo web application for tracking completion metrics and progress over time.

## Features

- **Completion Tracking**: Create and manage completion records with name, count, and timestamp
- **Full CRUD Operations**: Create, read, update, and delete completion entries
- **Responsive UI**: Bootstrap 5 based interface with navigation
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
- **Completions**: [http://127.0.0.1:3000/completions](http://127.0.0.1:3000/completions)
  - View all completion records
  - Create new completions
  - Edit existing completions
  - Delete completions

### API Endpoints
All completion endpoints support JSON and XML:
- `GET /completions` - List all completions
- `GET /completions/{id}` - Get specific completion
- `POST /completions` - Create new completion
- `PUT /completions/{id}` - Update completion
- `DELETE /completions/{id}` - Delete completion

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
