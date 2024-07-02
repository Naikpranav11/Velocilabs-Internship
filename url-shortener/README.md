# Velocilabs-Internship
This repo contains the test projects done during the internship period

# URL Shortener Application

This is a simple URL shortener application built using React for the frontend and Go (Golang) for the backend, with SQLite as the database.

## Features

- Shorten long URLs into shorter, more manageable links.
- Redirect users from shortened links to the original URLs.
- Store all URLs, whether shortened or not, in the SQLite database for future reference.

## Technologies Used

- **Frontend**: React
- **Backend**: Go (Golang)
- **Database**: SQLite

## Installation

### Backend (Go)

1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd url-shortener-backend

2. Dependencies:
    go get github.com/gorilla/sessions github.com/julienschmidt/httprouter github.com/justinas/alice modernc.org/sqlite golang.org/x/text github.com/davidmytton/url-verifier


## Directory Structure
.
├── data
├── internals
│   └── models
├── static
│   └── css
└── templates
