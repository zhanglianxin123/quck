#!/bin/bash
# 参数传入docker-compose.yaml路径
COMPOSE_FILE=$1

# 停止并删除之前的容器
docker-compose -f $COMPOSE_FILE down

# 重新创建和启动容器
docker-compose -f $COMPOSE_FILE up -d
