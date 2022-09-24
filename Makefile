run:
	go run cmd/server/main.go

db:
	docker-compose -f .docker/docker-compose.dev.yaml up --build -d

dbinit:
	mysql -h 127.0.0.1 -u root -p < tools/init/db.sql

goose:
ifeq ($(env), $(filter $(env),local test))
	db=onlinejudge; \
	goose -dir tools/migrations mysql root:oj-root-pass@tcp\(localhost:3306\)/$$db?parseTime=true $(c)
endif
