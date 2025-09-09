# go_crud_happens_backend
заметочник
1. Запуск
go run main.go db.go handlers.go models.go
2. Запросы

# Добавить заметку
curl -X POST -d '{"title":"Hello","content":"World"}' http://localhost:8080/notes -H "Content-Type: application/json"

# Получить все заметки
curl http://localhost:8080/notes

# Получить заметку по ID
curl http://localhost:8080/notes/1

# Поменять заметку по ID
curl -X PUT http://localhost:8080/notes/1 \
     -H "Content-Type: application/json" \
     -d '{"title": "Updated title", "content": "Updated content"}'

# Удалить заметку по ID
curl -X DELETE http://localhost:8080/notes/1