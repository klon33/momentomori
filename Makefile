
up: ## Запустить все сервисы (production)
	docker compose  up --build -d

down: ## Остановить все сервисы
	docker compose down

restart: ## Перезапустить все сервисы
	docker compose restart

logs: ## Показать логи всех сервисов
	docker compose logs -f
