#!/bin/bash

cd /opt/api/eas_api || exit

echo "📥 正在拉取最新代码..."
git fetch origin
git reset --hard origin/main

# 确保 DOCKER_HOST 不会干扰
unset DOCKER_HOST

echo "🐳 正在构建并启动 Docker 容器..."

# 使用 V2 命令：docker compose（中间有空格）
docker compose down --remove-orphans
docker compose up -d --build --force-recreate

echo "✅ 部署完成！"