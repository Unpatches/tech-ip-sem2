## Имя: Дорджиев Виктор
## Группа: ЭФМО-02-25
# ПЗ №1 — Микросервисы Auth + Tasks

## Цель
Научиться декомпозировать небольшую систему на два сервиса и
организовать корректное синхронное взаимодействие по HTTP (с
таймаутами, статусами ошибок и прокидыванием request-id)

В рамках ПЗ мы делаем учебную систему из двух компонентов:
1) Auth service — отвечает за “проверку доступа”
(упрощённая логика).
2) Tasks service — CRUD задач, но каждый запрос требует
проверки через Auth

## Установка и запуск

(Необходимы предустановленные Go версии 1.22 и выше и Git)

Клонировать репозиторий:

```
git clone <URL_РЕПОЗИТОРИЯ>
cd tech-ip-sem2
```

Команда запуска сервера:

Терминал 1
```
go run ./services/auth/cmd/auth
```
Терминал 2
```
go run ./services/tasks/cmd/tasks
```

## Структура проекта
```plaintext
tech-ip-sem2/
├── go.mod
├── go.sum
├── cmd/
│   ├── auth/
│   │   └── main.go
│   └── tasks/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── service/
│   │   │   └── auth.go
│   │   └── http/
│   │       ├── router.go
│   │       └── handlers/
│   │           └── auth.go
│   ├── tasks/
│   │   ├── service/
│   │   │   └── tasks.go
│   │   ├── client/
│   │   │   └── authclient.go
│   │   └── http/
│   │       ├── router.go
│   │       └── handlers/
│   │           └── tasks.go
│   └── shared/
│       ├── httpx/
│       │   └── json.go
│       └── middleware/
│           ├── logging.go
│           └── requestid.go
├── docs/
│   ├── pz1_api.md
│   └── pz1_diagram.md
├── README.md
└── .gitignore
```

## Границы ответственности
Auth service
* выдаёт “токен” (упрощённо),
* проверяет токен,
* возвращает информацию: валиден/не валиден.

Tasks service
* хранит и управляет задачами,
* перед выполнением операций проверяет токен через Auth.

## Схема взаимодействия
```mermaid
sequenceDiagram
    participant C as Client
    participant T as Tasks service
    participant A as Auth service
    C->>T: Request with Authorization
    T->>A: GET /v1/auth/verify (timeout 2-3s)

    alt Valid token
        A-->>T: 200 OK (valid)
        T-->>C: 200 OK / 201 Created / 204 No Content
    else Invalid token
        A-->>T: 401 Unauthorized
        T-->>C: 401 Unauthorized
    else Forbidden action
        A-->>T: 200 OK (valid)
        T-->>C: 403 Forbidden
    else Auth timeout or server error
        A-->>T: timeout / 5xx
        T-->>C: 502 Bad Gateway / 503 Service Unavailable
    end
```

## Список эндпоинтов (Auth и Tasks)

- Auth service
  - `POST /v1/auth/login`
  - `GET /v1/auth/verify`
- Tasks service
  - `POST /v1/tasks`
  - `GET /v1/tasks`
  - `GET /v1/tasks/{id}`
  - `PATCH /v1/tasks/{id}`
  - `DELETE /v1/tasks/{id}`


## Учётные данные для demo

- username: `student`
- password: `student`
- token: `demo-token`

## Скриншоты
### Скрин/лог с request-id, подтверждающий прокидывание.

<img width="1795" height="565" alt="image" src="https://github.com/user-attachments/assets/c4f2dfdc-aca1-489c-b1cb-574492ee860e" />


<img width="1796" height="513" alt="image" src="https://github.com/user-attachments/assets/23db8837-c6b1-4010-9c75-62ced7f8a078" />

### Получить токен

```
http://185.250.46.179:8081/v1/auth/login
```

<img width="479" height="461" alt="image" src="https://github.com/user-attachments/assets/60881aac-9dfa-4d68-8dcb-b2212fd19967" />

### Проверка токена напрямую

```
http://185.250.46.179:8081/v1/auth/verify
```

<img width="597" height="458" alt="image" src="https://github.com/user-attachments/assets/7f9939e4-980f-41c6-aeb7-c624bb15e4c5" />

### Создать задачу через Tasks (с проверкой Auth)

```
http://185.250.46.179:8082/v1/tasks
```

<img width="568" height="517" alt="image" src="https://github.com/user-attachments/assets/a3b9844e-81e8-400c-aaa2-65dc9feee181" />

### Попробовать без токена (должно быть 401)

```
http://185.250.46.179:8082/v1/tasks
```

<img width="536" height="455" alt="image" src="https://github.com/user-attachments/assets/bd55ad75-af30-4ac0-b412-51ddb6e47a78" />

