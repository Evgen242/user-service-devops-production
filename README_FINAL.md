# 🚀 User Service - Complete DevOps Project

![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-CI/CD-blue)
![Docker](https://img.shields.io/badge/Docker-Containerized-green)
![Go](https://img.shields.io/badge/Go-Microservice-success)
![Prometheus](https://img.shields.io/badge/Prometheus-Monitoring-orange)
![Grafana](https://img.shields.io/badge/Grafana-Dashboards-yellow)

## 📋 Project Overview
A complete DevOps-ready microservice with full monitoring stack, CI/CD, and production configuration.

## 🏗️ Architecture
User Service (Go) → PostgreSQL → Prometheus → Grafana
↓
Docker + Docker Compose
↓
GitHub Actions CI/CD

text

## 🚀 Quick Start
```bash
# Production deployment
docker-compose -f docker-compose-corrected.yml up -d

# Verify
curl http://localhost:8090/health
📊 Monitoring
Grafana: http://localhost:3000 (admin/admin123)

Prometheus: http://localhost:9090

Metrics: http://localhost:8090/metrics

🔧 Features
✅ Go REST API with Echo framework

✅ PostgreSQL with health checks

✅ Docker multi-stage builds

✅ Docker Compose orchestration

✅ Prometheus metrics collection

✅ Grafana dashboards

✅ GitHub Actions CI/CD

✅ Security scanning

✅ Production optimization

🌐 URLs
Repository: https://github.com/Evgen242/user-service-devops-production

CI/CD: https://github.com/Evgen242/user-service-devops-production/actions

📁 Project Structure
text
├── .github/workflows/    # CI/CD pipelines
├── cmd/api/             # Go application
├── pkg/                 # Internal packages
├── grafana/             # Dashboards & provisioning
├── docker-compose-*.yml # Environment configs
└── Dockerfile           # Multi-stage build
