# 📘 Цитатник — мини REST API на Go

Мини-сервис для хранения и управления цитатами. Проект разработан для демонстрации навыков работы с HTTP, организации кода, использования стандартной библиотеки Go и контейнеризации с помощью Docker.

---

## 🚀 Запуск проекта

Сервис запускается через `docker-compose`. Убедитесь, что у вас установлены Docker и docker-compose.

### 🔧 Сборка и запуск:

```bash
docker-compose up --build
```

После запуска API будет доступен по адресу:
http://localhost:8080

---

# 📌 Эндпоинты API

1. Добавление новой цитаты POST /quotes
  ```bash
curl -X POST http://localhost:8080/quotes \
  -H "Content-Type: application/json" \
  -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

2. Получение всех цитат GET /quotes
  ```bash
curl http://localhost:8080/quotes
```

3. Получение случайной цитаты GET /quotes/random
```bash
curl http://localhost:8080/quotes/random
```

4. Фильтрация по автору GET /quotes?author=AuthorName

Важно: из-за особенностей библиотеки gorilla/mux, используйте кавычки при вызове:
 ```bash
curl "http://localhost:8080/quotes?author=Confucius"
```

5. Удаление цитаты по ID DELETE /quotes/{id}
```bash
curl -X DELETE http://localhost:8080/quotes/1
```

---

# 🛠 Технические детали
- Язык: Go

- Фреймворк роутинга: gorilla/mux

- Хранилище данных: в памяти

- Контейнеризация: Docker + docker-compose

- Зависимости: только стандартная библиотека и gorilla/mux


