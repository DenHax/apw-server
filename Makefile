all:
	run-serv

compose-create-env:
	. ./scripts/create-env-compose.sh

compose-run:
	. ./scripts/docker_compose_run.sh

compose-down:
	. ./scripts/docker_compose_down.sh

compose-autostart:
	@$(MAKE) compose-run
	@$(MAKE) create-db

create-db:
	. ./scripts/create-db-psql.sh

look-psql:
	. ./scripts/lookup-storage.sh

run-serv:
	CONFIG_PATH=${CONFIG_PATH} . scripts/run-serv.sh

