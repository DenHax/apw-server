all:
	run-serv

activate-env:
	. ./scripts/activate-env.sh

compose-create-env:
	. ./scripts/env-compose.sh

compose-run:
	. ./scripts/docker_compose_run.sh

compose-down:
	. ./scripts/docker_compose_down.sh

create-db:
	. ./scripts/create-db-psql.sh

look-psql:
	. ./scripts/lookup-storage.sh

run-serv:
	CONFIG_PATH=${CONFIG_PATH} . scripts/run-serv.sh

auto-start:
	. scripts/autostart.sh
	. scripts/psql_start.sh
	@$(MAKE) run-serv

compose-autostart:
	@$(MAKE) compose-create-env
	@$(MAKE) activate-env ENV=compose
	@$(MAKE) compose-run
	@$(MAKE) create-db
