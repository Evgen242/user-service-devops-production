# 🚀 User Service - Production DevOps Project

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
| **User Service API** | [http://103.125.216.110:8090](http://103.125.216.110:8090) | ✅ Live | - |
| **Grafana Dashboards** | [http://103.125.216.110:3000](http://103.125.216.110:3000) | ✅ Live | `admin` / `admin123` |
| **Prometheus Metrics** | [http://103.125.216.110:9090](http://103.125.216.110:9090) | ✅ Live | - |
| **GitHub Repository** | [https://github.com/Evgen242/user-service-devops-production](https://github.com/Evgen242/user-service-devops-production) | 🔄 Active | - |
| **CI/CD Pipeline** | [https://github.com/Evgen242/user-service-devops-production/actions](https://github.com/Evgen242/user-service-devops-production/actions) | ✅ Working | - |

## 🏗️ System Architecture

cat > README.md << 'EOF'
# 🚀 User Service - Production DevOps Project

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
| **User Service API** | [http://103.125.216.110:8090](http://103.125.216.110:8090) | ✅ Live | - |
| **Grafana Dashboards** | [http://103.125.216.110:3000](http://103.125.216.110:3000) | ✅ Live | `admin` / `admin123` |
| **Prometheus Metrics** | [http://103.125.216.110:9090](http://103.125.216.110:9090) | ✅ Live | - |
| **GitHub Repository** | [https://github.com/Evgen242/user-service-devops-production](https://github.com/Evgen242/user-service-devops-production) | 🔄 Active | - |
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
