#!/bin/bash
# /opt/你的项目名/deploy.sh

cd /opt/api/eas_api || exit

echo "📥 正在拉取最新代码..."
git pull origin main

echo "🐳 正在构建并启动 Docker 容器..."
docker-compose down
docker-compose up -d --build

echo "✅ 部署完成！"