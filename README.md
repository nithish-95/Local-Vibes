# Local-Vibes

Local-Vibes is a full-stack web application designed for creating, discovering, and joining local events. It provides a platform for users to connect with others in their community through shared activities.

## Overview

This project consists of a Go-based backend that serves a RESTful API and a Vue.js single-page application for the frontend. The application allows users to register, log in, create their own events, and join events created by others.

## Features

### Event Creation and Management
- **Event Posting:** Users can create and post local events with details like title, description, date, time, location, and capacity.
- **Custom Rules:** Event creators can define and display custom rules alongside event details.
- **Event Dashboard:** A personalized dashboard for each user listing the events they are hosting and the events they have joined.
- **CRUD Operations:** Full Create, Read, Update, and Delete functionality for events created by the user.

### Event Discovery and Joining
- **Local Event Search:** Users can browse and search for local events available in their area.
- **Event Joining:** Users can easily join an event with a single click.
- **Leave Events:** Users can leave an event they have previously joined.

### User Management
- **Authentication:** Secure user registration and login functionality.
- **Session Management:** Persistent user sessions to keep users logged in.
- **User Privacy:** User data is kept private and is not shared with other users.
- **Pseudonymity:** Users can choose a username or alias to interact on the platform.

## Technology Stack

### Backend
- **Language:** Go
- **Web Framework/Router:** Chi
- **Database:** SQLite
- **Key Libraries:** `mattn/go-sqlite3`, `go-chi/chi`

### Frontend
- **Framework:** Vue.js (v3)
- **Routing:** Vue Router
- **Styling:** Tailwind CSS
- **Build Tool:** Vite
- **API Communication:** Fetch API

## Project Structure

The project is organized into two main parts: a `backend` and a `frontend`.

```
/
├── backend/
│   ├── cmd/            # Main application entry points
│   ├── internal/       # Internal application logic (database, models, routes, services)
│   └── ...
├── frontend/
│   ├── src/            # Vue.js application source code (components, views, router, api)
│   ├── public/         # Public assets
│   └── ...
├── database/           # Contains the SQLite database file (app.db)
├── .gitignore
├── LICENSE
└── README.md
```

## Getting Started

### Prerequisites

- Go (v1.18+ recommended)
- Node.js and npm

### Installation & Running

**Backend:**
```bash
cd backend/cmd/web
go run --tags "fts5" main.go
```
The backend server will start on `http://localhost:8080`.

**Frontend:**
```bash
cd frontend
npm install
npm run dev
```
The frontend development server will be accessible at `http://localhost:5173` (or another port if 5173 is busy).
