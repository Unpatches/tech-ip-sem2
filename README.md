# ПЗ №1 — Микросервисы Auth + Tasks

## Что реализовано

- Auth service
  - `POST /v1/auth/login`
  - `GET /v1/auth/verify`
- Tasks service
  - `POST /v1/tasks`
  - `GET /v1/tasks`
  - `GET /v1/tasks/{id}`
  - `PATCH /v1/tasks/{id}`
  - `DELETE /v1/tasks/{id}`
- Проверка токена через межсервисный HTTP-вызов
- Timeout 3 секунды
- Прокидывание `X-Request-ID`
- Логирование

## Учётные данные для demo

- username: `student`
- password: `student`
- token: `demo-token`

## Запуск

### Auth service

```bash
cd services/auth
export AUTH_PORT=8081
cd ../..
go run ./services/auth/cmd/auth