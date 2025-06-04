.PHONY: dbinit goose load-env

load-env:
ifneq (,$(wildcard .env.local))
	$(eval include .env.local)
	$(foreach var,$(shell cat .env.local 2>/dev/null | grep -v ^\# | cut -d= -f1),$(eval export $(var)))
endif
ifneq (,$(wildcard .env))
	$(eval include .env)
	$(foreach var,$(shell cat .env 2>/dev/null | grep -v ^\# | cut -d= -f1),$(eval export $(var)))
endif

dbinit: load-env
	psql "postgresql://$(WRITER_DB_USER):$(WRITER_DB_PASS)@$(WRITER_DB_HOST):$(WRITER_DB_PORT)/$(WRITER_DB_NAME)" -f tools/init/db.sql

goose: load-env
ifeq ($(env), $(filter $(env),local test))
	goose -dir tools/migrations postgres "postgresql://$(WRITER_DB_USER):$(WRITER_DB_PASS)@$(WRITER_DB_HOST):$(WRITER_DB_PORT)/$(WRITER_DB_NAME)" $(c)
endif
