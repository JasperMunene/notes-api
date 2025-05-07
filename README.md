# 📝 Notes API

A simple RESTful API built with Go for creating, retrieving, updating, and deleting notes.  
It uses PostgreSQL as the backend database and is structured with clear separation of concerns (handlers, db logic, and models).

---

## 🚀 Features

- `GET /notes` - Retrieve all notes
- `POST /notes` - Create a new note
- `PUT /notes?id=ID` - Update an existing note by ID
- `DELETE /notes?id=ID` - Delete a note by ID
- JSON request and response format
- Proper error handling
- Environment-based configuration (via `.env`)
- Clean project structure (handlers, db, models)

---

## 📁 Project Structure

```

notes-api/
│
├── main.go              # Entry point
├── handlers/            # HTTP request handlers
│   └── note.go
├── db/                  # Database connection and queries
│   └── db.go
├── models/              # Data models (e.g., Note struct)
│   └── note.go
├── .env                 # Environment variables (DATABASE\_URL)
└── go.mod / go.sum      # Go modules and dependencies

````

---

## 🛠️ Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/JasperMunene/notes-api.git
cd notes-api
````

### 2. Set Up the Database

Create a PostgreSQL database and table:

```sql
CREATE TABLE notes (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  created TIMESTAMP NOT NULL
);
```

### 3. Configure Environment Variables

Create a `.env` file in the project root:

```
DATABASE_URL=postgres://username:password@localhost:5432/your_db
```

### 4. Run the Server

```bash
go run main.go
```

Server will start on:

```
http://localhost:8080
```

---

## 📦 API Endpoints

### `GET /notes`

**Description**: Returns a list of notes, ordered by creation time (descending).

**Response**:

```json
[
  {
    "id": 1,
    "title": "First Note",
    "content": "This is the first note.",
    "created_at": "2024-08-04T10:10:10Z"
  }
]
```

---

### `POST /notes`

**Description**: Create a new note.

**Request Body**:

```json
{
  "title": "New Note",
  "content": "This is a new note."
}
```

**Response**:

```json
{
  "id": 2,
  "title": "New Note",
  "content": "This is a new note.",
  "created_at": "2024-08-04T10:15:00Z"
}
```

---

### `PUT /notes?id=1`

**Description**: Update a note by ID.

**Request Body**:

```json
{
  "title": "Updated Title",
  "content": "Updated content goes here."
}
```

**Response**:

```json
{
  "id": 1,
  "title": "Updated Title",
  "content": "Updated content goes here.",
  "created_at": "2024-08-04T10:10:10Z"
}
```

---

### `DELETE /notes?id=1`

**Description**: Delete a note by ID.

**Response**:

```json
{
  "message": "Note deleted successfully"
}
```

---

## 📚 Tech Stack

* **Language**: Go (Golang)
* **Database**: PostgreSQL
* **Standard Library**: `net/http`, `database/sql`, `encoding/json`
* **Driver**: `lib/pq` for PostgreSQL
* **Env Loader**: `joho/godotenv`

---

## 🧼 Code Style

* Idiomatic Go structure (handlers, db access, models)
* GoDoc comments for all exported symbols
* Error logging via standard `log` package
* Clean separation of concerns

---

## 🧪 Future Improvements

* Authentication (JWT or session-based)
* Pagination for `GET /notes`
* URL path routing using `chi` or `gorilla/mux` for cleaner endpoints

---

## 📄 License

MIT License

---

## 👨‍💻 Author

**Jasper Munene**
[devjaspermunene@gmail.com](mailto:devjaspermunene@gmail.com)
