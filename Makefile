.PHONY: help mysql-up mysql-down mysql-logs mysql-wait run start stop

COMPOSE := docker compose -f docker-compose.yml
MYSQL_ROOT_PASSWORD := root123

help:
	@echo "常用命令:"
	@echo "  make start      - 启动 MySQL 容器并运行项目（推荐）"
	@echo "  make mysql-up   - 只启动 MySQL 容器"
	@echo "  make mysql-down - 停止并移除 MySQL 容器"
	@echo "  make mysql-logs - 查看 MySQL 日志"
	@echo "  make run        - 只运行 Go 项目（需 MySQL 已启动）"
	@echo "  make stop       - 停止 MySQL 容器"

mysql-up:
	$(COMPOSE) up -d mysql
	@$(MAKE) mysql-wait

mysql-wait:
	@echo "等待 MySQL 就绪..."
	@i=0; \
	until $(COMPOSE) exec -T mysql mysqladmin ping -h 127.0.0.1 -uroot -p$(MYSQL_ROOT_PASSWORD) --silent 2>/dev/null; do \
		i=$$((i+1)); \
		if [ $$i -gt 30 ]; then \
			echo "MySQL 启动超时，请检查: make mysql-logs"; \
			exit 1; \
		fi; \
		sleep 2; \
	done
	@echo "MySQL 已就绪"

mysql-down:
	$(COMPOSE) down

mysql-logs:
	$(COMPOSE) logs -f mysql

run:
	go run .

start: mysql-up
	go run .

stop: mysql-down
