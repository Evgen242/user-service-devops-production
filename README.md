k# User Service - DevOps Test Project

## Описание
Микросервис для управления пользователями с полным DevOps стеком.

## Технологии
- Go + Echo framework
- PostgreSQL
- Docker + Docker Compose
- Prometheus + Grafana
- GitHub Actions CI/CD

## Быстрый старт
```bash
docker-compose up -d
API Endpoints
GET /health - Health check

GET /api/users - Получить пользователей

POST /api/users - Создать пользователя

CI/CD
Проект использует GitHub Actions 
