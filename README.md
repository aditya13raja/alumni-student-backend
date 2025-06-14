# 🎓 Alumni-Student Backend

This is the **backend service** for the Alumni-Student Platform — a web app that connects alumni and students for blogging, chatting, job sharing, and discussion forums.

Built with **Golang**, **Fiber**, and **MongoDB**, the backend handles user authentication, RESTful APIs for core features, and real-time messaging using **Pusher**.

---

## 🔗 Frontend repo link
**[Frontend Repository](https://github.com/aditya13raja/alumni-student-frontend)**

---

## 🚀 Features

- 🔐 JWT-based authentication with middleware
- 💬 Real-time chat with Pusher
- 📚 Topic-wise discussion threads
- ✍️ Blog system (create/read)
- 💼 Alumni-only job postings
- 🧾 Category-based content organization
- 📦 Modular route/controller structure

---

## 🛠️ Tech Stac

| Layer       | Tech                 |
|-------------|----------------------|
| Language    | Go (Golang)          |
| Framework   | Fiber (Express-like) |
| Database    | MongoDB              |
| Real-time   | Pusher Channels      |
| Auth        | JWT (Token-based)    |

---

## 📁 Project Structure

```
├── controllers/ # Core business logic  
├── models/ # Data models (MongoDB schemas)  
├── routes/ # Grouped route definitions  
├── utils/ # Helper functions and DB clients  
├── configs/ # External service setup (e.g., Pusher)  
├── main.go # App entry point
```

---

## 📦 API Modules

### ✅ Auth
- `POST /api/auth/login` — Login user and return token
- `POST /api/auth/register` — Register new user

### 💬 Chat
- `POST /api/chat/send` — Send message (Pusher broadcast)
- `GET /api/chat/get/:topic` — Get messages by topic

### 📚 Topics
- `POST /api/topics/create-topic` — Create a topic
- `GET /api/topics/get-topics` — Get all topics
- `GET /api/topics/` — Get topics by category

### 🧾 Categories
- `POST /api/category/create` — Add new category
- `GET /api/category/get` — Get all categories

### ✍️ Blog
- `POST /api/blog/save-blog` — Save blog (Alumni only)
- `GET /api/blog/:id` — Fetch blog by ID
- `GET /api/blog/list/blogs` — List all blogs
- `GET /api/blog/latest/blogs` — Latest 5 blogs

### 💼 Jobs
- `POST /api/jobs/create` — Post job (Alumni only)
- `GET /api/jobs/:id` — Get job by ID
- `GET /api/jobs/list/jobs` — List all jobs
- `GET /api/jobs/latest/jobs` — Latest 6 jobs

---

## 🔧 Setup Instructions

### Prerequisites
- Go 1.19+
- MongoDB running locally or on Atlas
- Pusher App credentials

### 1. Clone the Repo
```bash
git clone https://github.com/your-username/alumni-student-backend.git
cd alumni-student-backend
````

### 2. Set up `.env`

```env
PORT=1234
MONGO_URI=your_mongodb_connection_string
JWT_SECRET=your_secret_key
PUSHER_APP_ID=your_app_id
PUSHER_KEY=your_key
PUSHER_SECRET=your_secret
PUSHER_CLUSTER=your_cluster
```

### 3. Run the Server

```bash
go mod tidy
go run main.go
```

---

## 🧪 Testing the APIs

You can test endpoints using:

- Postman (import collection)

- cURL

- Frontend app

Authentication-protected routes require a valid JWT in the `Authorization` header.

---

## 👥 Contributing

PRs and issues are welcome! Please include tests and maintain consistent formatting.

---
