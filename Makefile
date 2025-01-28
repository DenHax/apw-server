all:
	run-serv

run-serv:
	CONFIG_PATH=${CONFIG_PATH} . script/run-serv.sh

compose-run:
	. ./script/docker_compose_run.sh

compose-down:
	. ./script/docker_compose_down.sh

create-db:
	. ./script/create-db-psql.sh

look-psql:
	. ./script/lookup-storage.sh

auto-start:
	. script/autostart.sh
	. script/psql_start.sh
	@$(MAKE) run-serv

compose-autostart:
	. ./script/env-compose.sh
	ENV=compose . ./script/activate-env.sh
	@$(MAKE) compose-run
	. ./script/create-db-psql.sh
