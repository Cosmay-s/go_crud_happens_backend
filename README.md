# Go CRUD Заметочник с SQLite

Простой REST API для заметок на Go с использованием SQLite (пакет `glebarez/sqlite`).

## Функционал

- Создание заметки — `POST /notes`
- Получение списка всех заметок — `GET /notes`
- Получение заметки по ID — `GET /notes/{id}`
- Обновление заметки по ID — `PUT /notes/{id}`
- Удаление заметки по ID — `DELETE /notes/{id}`

## Запуск проекта

1. Убедитесь, что у вас установлен Go 1.18 или выше.
2. Установите зависимости:
   ```bash
   go get github.com/glebarez/sqlite
   ```
3. Запустите проект:
   ```bash
   go run .
   ```
   Или укажите все файлы явно:
   ```bash
   go run main.go db.go model.go handlers.go
   ```

## API команды и примеры

### Создать заметку
```bash
curl -X POST http://localhost:8080/notes \
     -H "Content-Type: application/json" \
     -d '{"title": "Заголовок", "content": "Текст заметки"}'
```

### Получить все заметки
```bash
curl http://localhost:8080/notes
```

### Получить заметку по ID
```bash
curl http://localhost:8080/notes/1
```

### Обновить заметку по ID
```bash
curl -X PUT http://localhost:8080/notes/1 \
     -H "Content-Type: application/json" \
     -d '{"title": "Новый заголовок", "content": "Обновлённый текст"}'
```

### Удалить заметку по ID
```bash
curl -X DELETE http://localhost:8080/notes/1
```

## Структура проекта

- `main.go` — запуск сервера и настройка маршрутов
- `db.go` — подключение и инициализация базы данных
- `model.go` — структура заметки
- `handlers.go` — HTTP-обработчики и функции для работы с БД

## Зависимости

- Go 1.18+
- `github.com/glebarez/sqlite` — драйвер SQLite