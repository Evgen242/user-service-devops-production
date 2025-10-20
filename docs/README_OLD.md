# 🚀 User Service - Production DevOps Project

## 🌐 LIVE PRODUCTION URLs
- **User Service API**: http://103.125.216.110:8090
- **Grafana Dashboards**: http://103.125.216.110:3000 (admin/admin123) 
- **Prometheus Metrics**: http://103.125.216.110:9090
- **GitHub Repository**: https://github.com/Evgen242/user-service-devops-production
- **CI/CD Pipeline**: https://github.com/Evgen242/user-service-devops-production/actions

## 🏗️ Production Architecture
Internet Users
↓
http://103.125.216.110:8090 (User Service API)
↓
http://103.125.216.110:5433 (PostgreSQL Database)
↓
http://103.125.216.110:9090 (Prometheus Monitoring)
↓
http://103.125.216.110:3000 (Grafana Dashboards)

text

## 🚀 Quick Start
```bash
# Deploy on production server
docker-compose -f docker-compose-corrected.yml up -d

# Test from anywhere in the world
curl http://103.125.216.110:8090/health
curl http://103.125.216.110:8090/api/users
📊 Live Monitoring Stack
API Health: http://103.125.216.110:8090/health

Metrics: http://103.125.216.110:8090/metrics

Prometheus UI: http://103.125.216.110:9090

Grafana: http://103.125.216.110:3000

✅ Production Features
Public REST API accessible globally

Real-time metrics and monitoring

Automated CI/CD deployment

Health checks and observability

Docker containerization

PostgreSQL persistence
