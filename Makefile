all:
	run-serv

run-serv:
	CONFIG_PATH=${CONFIG_PATH} . script/run-serv.sh

run-compose:
	CONFIG_PATH=${CONFIG_PATH} . script/deploy-start-compose.sh

down-compose:
	. ./script/docker_compose_down.sh

auto-start:
	. script/autostart.sh
	. script/psql_start.sh
	@$(MAKE) run-serv

auto-start-compose:
	. script/autostart-compose.sh
	@$(MAKE) run-compose
