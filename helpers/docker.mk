# ================================================
# DOCKER SHORTCUTS
# ================================================

.PHONY: verify-docker
verify-docker:
ifeq (, $(shell which docker))
	$(error "No docker in $(PATH), consider installing it from https://docs.docker.com/install")
endif
	docker --version

.PHONY: docker-stop
docker-stop:
	$(if $(strip $(DOCKER_CONTAINER_LIST)), docker stop $(DOCKER_CONTAINER_LIST))

.PHONY: docker-remove
docker-remove:
	$(if $(strip $(DOCKER_CONTAINER_LIST)), docker rm $(DOCKER_CONTAINER_LIST))

.PHONY: docker-volume-prune
docker-volume-prune:
	@-docker volume prune -f

.PHONY: docker-network-prune
docker-network-prune:
	@-docker network prune -f

.PHONY: docker-system-prune
docker-system-prune:
	@-docker system prune -af

.PHONY: docker-clean
docker-clean: verify-docker docker-stop docker-remove docker-volume-prune docker-network-prune
	@echo "# --------------------------------------"
	@echo "# Cleaning Docker Environment"
	@echo "# --------------------------------------"

.PHONY: docker-deep-clean
docker-deep-clean: verify-docker docker-stop docker-remove docker-volume-prune docker-network-prune docker-system-prune
	@echo "# --------------------------------------"
	@echo "# Deep cleaning Docker Environment"
	@echo "# --------------------------------------"

# ================================================
# DOCKER COMPOSE SHORTCUTS
# ================================================

.PHONY: compose-up
compose-up:
	docker-compose up --build

.PHONY: compose-up-background
compose-up-background:
	docker-compose up --build -d

.PHONY: compose-down
compose-down:
	docker-compose down

.PHONY: compose-ps
compose-ps:
	docker-compose ps

.PHONY: compose-run
compose-run: compose-ps compose-down compose-up
ifneq ("$(wildcard $(./.env))","")
  compose-ps compose-down compose-up
endif

.PHONY: compose-run-background
compose-run-background: compose-ps compose-down compose-up-background
ifneq ("$(wildcard $(./.env))","")
  compose-ps compose-down compose-up-background
endif
