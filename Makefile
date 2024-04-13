.PHONY: up down

PROJECT_NAME ?= tinyurl

up:
	docker compose -p ${PROJECT_NAME} up

down:
	docker compose down \
		--rmi all \
		--volumes \
		--remove-orphans
