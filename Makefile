run:
	go run cmd/server/main.go

localdb:
	docker-compose -f .docker/docker-compose.dev.yaml up --build -d

goose:
ifeq ($(env), $(filter $(env),local test))
	db=onlinejudge; \
	goose -dir tools/migrations mysql judge_admin:judge_pass@tcp\(localhost:3306\)/$$db?parseTime=true $(c)
endif
