#!/bin/bash
echo "🧪 PRODUCTION SERVICE TESTS - 103.125.216.110"
echo "============================================="

SERVER="103.125.216.110"

echo "1. Testing User Service API..."
response=$(curl -s -w "%{http_code}" http://$SERVER:8090/health)
status=$(echo "$response" | tail -n1)
if [ "$status" -eq 200 ]; then
echo " ✅ HTTP 200 - Service Healthy"
echo " 📍 URL: http://$SERVER:8090/health"
else
echo " ❌ HTTP $status - Service Issues"
fi

echo "2. Testing Prometheus..."
if curl -s http://$SERVER:9090/-/healthy > /dev/null; then
echo " ✅ Prometheus Healthy"
echo " 📍 URL: http://$SERVER:9090"
else
echo " ❌ Prometheus Unhealthy"
fi

echo "3. Testing Grafana..."
grafana_status=$(curl -s http://$SERVER:3000/api/health | jq -r '.database' 2>/dev/null)
if [ "$grafana_status" = "ok" ]; then
echo " ✅ Grafana Healthy"
echo " 📍 URL: http://$SERVER:3000"
else
echo " ❌ Grafana Issues"
fi

echo ""
echo "🎯 ALL SERVICES ACCESSIBLE FROM:"
echo " Any computer → http://103.125.216.110"
echo ""
echo "💡 TIP: Share these URLs in your portfolio!"
