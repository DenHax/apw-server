all:
	run-serv

activate-env:
	. ./script/activate-env.sh

compose-create-env:
	. ./script/env-compose.sh

compose-run:
	. ./script/docker_compose_run.sh

compose-down:
	. ./script/docker_compose_down.sh

create-db:
	. ./script/create-db-psql.sh

look-psql:
	. ./script/lookup-storage.sh

run-serv:
	CONFIG_PATH=${CONFIG_PATH} . script/run-serv.sh

auto-start:
	. script/autostart.sh
	. script/psql_start.sh
	@$(MAKE) run-serv

compose-autostart:
	@$(MAKE) compose-create-env
	@$(MAKE) activate-env ENV=compose
	@$(MAKE) compose-run
	@$(MAKE) create-db
