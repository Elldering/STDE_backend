# STDE Backend

Высокопроизводительный REST API-сервер для мобильного приложения **STDE**.  
Написан на Go, использует PostgreSQL для хранения данных, Redis для кэша и AWS S3/MinIO для файлов. Контейнеризован с Docker, описан инфраструктурой в `docker-compose.yml`.

---

## Стек

| Слой          | Технологии |
|---------------|------------|
| Язык          | Go 1.22 +  |
| HTTP-framework| [Gin](https://github.com/gin-gonic/gin) |
| ORM / миграции| GORM & golang-migrate |
| Auth          | JWT (HS256) |
| БД            | PostgreSQL 15 |
| Кэш / очередь | Redis 7     |
| Объектное хранилище | AWS S3 / MinIO |
| Документация  | Swagger (OpenAPI 3) |
| Сборка/деплой | Docker + docker-compose, GitLab CI |

---

## Структура каталога

```text
cmd/                # точки входа (main.go)
internal/
  api/              # контроллеры, DTO
  service/          # бизнес-логика
  repository/       # доступ к данным (GORM)
  middleware/       # Auth, CORS, Logger
  cron/             # фоновые задачи
pkg/                # переиспользуемые пакеты
config/             # конфигурация и примеры .env
migrations/         # SQL-миграции
scripts/            # утилиты (lint, test, ci)
docs/               # Swagger yaml, diagrams
