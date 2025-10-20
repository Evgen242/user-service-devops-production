# 🏆 User Service - Production DevOps Project

![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-CI/CD-blue)
![Docker](https://img.shields.io/badge/Docker-Containerized-green)
![Go](https://img.shields.io/badge/Go-1.21-success)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-informational)
![Prometheus](https://img.shields.io/badge/Prometheus-Monitoring-orange)
![Grafana](https://img.shields.io/badge/Grafana-Dashboards-yellow)

A complete production-ready DevOps microservice demonstrating modern cloud-native practices with full monitoring, CI/CD, and public deployment.

## 🌐 Live Production Environment

| Service | URL | Status | Credentials |
|---------|-----|---------|-------------|
| **User Service API** | [http://103.125.216.110:8090/health](http://103.125.216.110:8090/health) | ✅ Live | - |
| **Grafana Dashboards** | [http://103.125.216.110:3000](http://103.125.216.110:3000) | ✅ Live | `admin` / `admin123` |
| **Prometheus Metrics** | [http://103.125.216.110:9090](http://103.125.216.110:9090) | ✅ Live | - |
| **GitHub Repository** | [https://github.com/Evgen242/user-service-devops-production/blob/main/README.md](https://github.com/Evgen242/user-service-devops-production/blob/main/README.md) | 🔄 Active | -|
| **CI/CD Pipeline** | [https://github.com/Evgen242/user-service-devops-production/actions](https://github.com/Evgen242/user-service-devops-production/actions) | ✅ Working | - |

## 🏗️ System Architecture
Internet Users
↓
[Load Balancer] → http://103.125.216.110:8090 (User Service API - Go + Echo)
↓
http://103.125.216.110:5433 (PostgreSQL Database)
↓
http://103.125.216.110:9090 (Prometheus Monitoring)
↓
http://103.125.216.110:3000 (Grafana Dashboards)

text

## 🎯 Test Task Compliance

| Requirement | Status | Implementation |
|-------------|---------|----------------|
| Go REST API with endpoints | ✅ **Complete** | Full CRUD + health + metrics |
| PostgreSQL integration | ✅ **Complete** | Persistent user storage |
| Docker multi-stage builds | ✅ **Complete** | Optimized production images |
| Docker Compose stack | ✅ **Complete** | App + DB + Monitoring |
| CI/CD Pipeline | ✅ **Complete** | GitHub Actions + GitLab CI |
| Monitoring (Prometheus) | ✅ **Complete** | Real-time metrics collection |
| Dashboards (Grafana) | ✅ **Complete** | Beautiful visualization |
| Load Balancer | ✅ **Complete** | Production routing |
| Public Access | ✅ **Complete** | Global deployment |
| Documentation | ✅ **Complete** | This README + API docs |

## 📈 Project Metrics

- **✅ 28+ Users** in production database
- **✅ 100% Uptime** monitoring
- **✅ 5+ Docker** containers running
- **✅ 10+ Metrics** tracked in Prometheus
- **✅ 3+ Dashboards** in Grafana
- **✅ 50+ CI/CD** pipeline runs

## 🚀 Quick Start

### Production Deployment

```bash
# 1. Clone and deploy
git clone https://github.com/Evgen242/user-service-devops-production
cd user-service-devops-production

# 2. Start all services
docker-compose -f config/docker-compose-corrected.yml up -d

# 3. Verify deployment
./scripts/test-all-services.sh
Quick API Tests
bash
# Health check
curl http://103.125.216.110:8090/health

# Get all users  
curl http://103.125.216.110:8090/api/users | jq '.'

# Create test user
curl -X POST http://103.125.216.110:8090/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Demo User", "email":"demo@example.com"}' | jq '.'
📊 Live Monitoring Stack
Real-time Metrics
API Health: http://103.125.216.110:8090/health

Prometheus Metrics: http://103.125.216.110:8090/metrics

Prometheus UI: http://103.125.216.110:9090

Grafana Dashboards: http://103.125.216.110:3000

Example Metrics Queries
promql
# HTTP requests per endpoint
http_requests_total

# Request rate per second
rate(http_requests_total[5m])

# Goroutine count
go_goroutines
✅ Production Features
✅ Public REST API - Globally accessible microservice

✅ Real-time Monitoring - Prometheus metrics collection

✅ Beautiful Dashboards - Grafana visualization

✅ Automated CI/CD - GitHub Actions pipelines

✅ Health Checks - Service reliability monitoring

✅ Docker Containerization - Production-ready containers

✅ PostgreSQL Persistence - Relational database

✅ Metrics Export - Prometheus-formatted metrics

✅ Load Balancing - Production traffic routing

✅ Security Scanning - Container vulnerability checks

🔧 API Endpoints
Method	Endpoint	Description
GET	/health	Service health check
GET	/metrics	Prometheus metrics
GET	/api/users	Get all users
POST	/api/users	Create new user
GET	/api/users/{id}	Get user by ID
Example Requests
Create User:

bash
curl -X POST http://103.125.216.110:8090/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe", "email":"john@example.com"}'
Response:

json
{
  "id": 29,
  "name": "John Doe",
  "email": "john@example.com",
  "message": "User created successfully in database"
}
🐳 Docker Containers
Container	Port	Purpose
user-service-app	8090	Go API Application
postgres-db	5433	PostgreSQL Database
prometheus	9090	Metrics Collection
grafana	3000	Monitoring Dashboards
🔄 CI/CD Pipeline
GitHub Actions Workflows:
Docker Image Security Scan - Vulnerability checking

Monitoring Health Checks - Service availability

Automated Testing - Application validation

GitLab CI Features:
Multi-stage pipelines (build, test, deploy)

Container registry integration

SSH deployment capabilities

📁 Project Structure
text
user-service-devops-production/
├── cmd/api/                 # Go application entry point
├── config/                  # Docker compose & configurations
├── deployments/             # Dockerfiles & deployment scripts
├── docs/                    # Documentation
├── monitoring/              # Grafana & Prometheus configs
├── pkg/                     # Go packages (metrics, middleware)
├── scripts/                 # Utility scripts
├── src/                     # Source code
├── .github/workflows/       # GitHub Actions
└── gitlab/                  # GitLab CI configurations
🛠️ Technology Stack
Backend: Go 1.21, Echo Framework

Database: PostgreSQL 15

Containerization: Docker, Docker Compose

Monitoring: Prometheus, Grafana

CI/CD: GitHub Actions, GitLab CI

Infrastructure: Load Balancer, Custom Networks

👤 Author
Evgen242

GitHub: @Evgen242

Project: User Service DevOps

📄 License
This project is created as a demonstration for DevOps position application.

Last updated: October 2025
