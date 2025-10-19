#!/bin/bash
set -e

echo "🚀 Starting deployment process..."

ENVIRONMENT=${1:-staging}
COMMIT_SHA=${2:-latest}

case $ENVIRONMENT in
  staging)
    echo "📦 Deploying to STAGING..."
    docker-compose -f docker-compose-staging.yml pull
    docker-compose -f docker-compose-staging.yml up -d
    ;;
  production)
    echo "🎯 Deploying to PRODUCTION..."
    docker-compose -f docker-compose-production.yml pull  
    docker-compose -f docker-compose-production.yml up -d
    ;;
  *)
    echo "❌ Unknown environment: $ENVIRONMENT"
    exit 1
    ;;
esac

echo "✅ Deployment completed successfully!"
echo "📊 Health check..."
sleep 10
curl -f http://localhost:8090/health || exit 1
