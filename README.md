# User Service - DevOps Demo

Микросервис для управления пользователями с полным CI/CD пайплайном.

## 🚀 Features

- REST API на Go с Echo framework
- PostgreSQL database
- Docker контейнеризация
- Health checks и метрики
- CI/CD с GitLab
- Мониторинг с Prometheus/Grafana

## 📚 API Endpoints

### Health Check
```http
GET /health
Проверка статуса приложения и БД.

Metrics
http
GET /metrics
Метрики для Prometheus.

Get Users
http
GET /api/users
Получить список всех пользователей.

Create User
http
POST /api/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com"
}
🐳 Local Development
bash
# Запуск с Docker Compose
docker-compose -f docker-compose-fixed.yml up -d --build

# Тестирование
curl http://localhost:8095/health
curl http://localhost:8095/api/users
🏗️ Architecture
Go 1.21 - основной язык

Echo framework - HTTP router

PostgreSQL 15 - база данных

Docker - контейнеризация

Multi-stage builds - оптимизация образов

🔧 Environment Variables
Variable	Description	Default
DB_HOST	Database host	postgres
DB_PORT	Database port	5432
DB_NAME	Database name	userdb
DB_USER	Database user	postgres
DB_PASSWORD	Database password	password
PORT	Application port	8090
📊 Monitoring
Application metrics via /metrics

Health checks via /health

Docker container health checks

PostgreSQL connection monitoring

## 🚀 **Теперь настраиваем GitLab CI/CD**

### 1. **Создаем .gitlab-ci.yml**
```bash
cat > .gitlab-ci.yml << 'EOF'
stages:
  - test
  - build
  - security-scan
  - deploy

variables:
  APP_NAME: "user-service"
  DOCKER_HOST: "tcp://docker:2375"

before_script:
  - docker info

unit-test:
  stage: test
  image: golang:1.21-alpine
  script:
    - go mod download
    - go test -v ./...
    - go vet ./...
  only:
    - merge_requests
    - main

build:
  stage: build
  image: docker:latest
  services:
    - docker:dind
  variables:
    DOCKER_TLS_CERTDIR: ""
  script:
    - docker build -t $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA .
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
  only:
    - main
    - develop

security-scan:
  stage: security-scan
  image: docker:latest
  services:
    - docker:dind
  script:
    - apk add --no-cache curl
    - curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/main/contrib/install.sh | sh -s -- -b /usr/local/bin
    - trivy image $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
  allow_failure: true
  only:
    - main

deploy-staging:
  stage: deploy
  image: alpine:latest
  before_script:
    - apk add --no-cache openssh-client
    - eval $(ssh-agent -s)
    - echo "$SSH_PRIVATE_KEY" | tr -d '\r' | ssh-add -
    - mkdir -p ~/.ssh
    - chmod 700 ~/.ssh
  script:
    - ssh -o StrictHostKeyChecking=no $DEPLOY_USER@$DEPLOY_SERVER "docker pull $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA"
    - ssh -o StrictHostKeyChecking=no $DEPLOY_USER@$DEPLOY_SERVER "docker-compose -f /opt/user-service/docker-compose.yml up -d"
  environment:
    name: staging
    url: http://staging.example.com
  only:
    - develop

deploy-production:
  stage: deploy
  image: alpine:latest
  before_script:
    - apk add --no-cache openssh-client
    - eval $(ssh-agent -s)
    - echo "$SSH_PRIVATE_KEY" | tr -d '\r' | ssh-add -
    - mkdir -p ~/.ssh
    - chmod 700 ~/.ssh
  script:
    - ssh -o StrictHostKeyChecking=no $DEPLOY_USER@$DEPLOY_SERVER "docker pull $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA"
    - ssh -o StrictHostKeyChecking=no $DEPLOY_USER@$DEPLOY_SERVER "docker-compose -f /opt/user-service/docker-compose.yml up -d"
  environment:
    name: production
    url: http://production.example.com
  only:
    - main
  when: manual
