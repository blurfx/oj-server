docker-up:
	docker-compose -f .docker/docker-compose.dev.yaml up --build -d
docker-down:
	docker-compose -f .docker/docker-compose.dev.yaml down

dbinit:
	psql "postgresql://judge_admin:judge_pass@localhost:5432/postgres" -f server/tools/init/db.sql

goose:
ifeq ($(env), $(filter $(env),local test))
	db=onlinejudge; \
	goose -dir server/tools/migrations postgres "postgresql://judge_admin:judge_pass@localhost:5432/$$db" $(c)
endif
