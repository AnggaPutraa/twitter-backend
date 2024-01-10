# 🔥 Twitter Backend
Just a simple twitter backend clone made with Go.
For full documentation regards API can be found [here](https://documenter.getpostman.com/view/31936842/2s9YsKhXoq).

## 🚀 Features
- JWT Authentication
- Following Other User
- Post a tweet
- Comment a tweet and replies other comment freely

## ⚒️ Setup Development Environment
### 📦 Required tools
- [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [sqlc](https://github.com/sqlc-dev/sqlc#installation)
- [air](https://github.com/cosmtrek/air)

### 🏗️ How to Generate Code
Download packages
```bash
make install
```
Generate development database
```bash
make db-up
```
Generate schemas
```bash
make migrate-up
```
Generate type-safe queries
```bash
make sqlc
```
Setup .env
```bash
DB_HOST=
DB_PORT=
DB_USERNAME=
DB_PASSWORD=
DB_DATABASE=-db
ACCESS_TOKEN_SECRET=
REFRESH_TOKEN_SECRET=
SERVER_ADDRESS=
```

### 🛫 How to Run Server
```bash
air
```