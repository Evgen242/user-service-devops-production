📚 Complete Project Documentation
🌐 Live Production URLs
User Service API: http://103.125.216.110:8090

Grafana Dashboards: http://103.125.216.110:3000 (admin/admin123)

Prometheus Metrics: http://103.125.216.110:9090

GitHub Repository: https://github.com/Evgen242/user-service-devops-production

🏗️ System Architecture
Internet Users
    ↓
[Load Balancer] → http://103.125.216.110:8090 (User Service API - Go + Echo)
    ↓
http://103.125.216.110:5433 (PostgreSQL Database)
    ↓  
http://103.125.216.110:9090 (Prometheus Monitoring)
    ↓
http://103.125.216.110:3000 (Grafana Dashboards)
🚀 Quick Start
# Deploy on production server
docker-compose -f config/docker-compose-corrected.yml up -d

# Test from anywhere in the world
curl http://103.125.216.110:8090/health
curl http://103.125.216.110:8090/api/users
📊 API Endpoints
GET /health - Service health check

GET /api/users - Get all users

POST /api/users - Create new user

GET /metrics - Prometheus metrics

⬆️ Back to Main README
