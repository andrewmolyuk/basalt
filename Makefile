init:
	@chmod +x ./scripts/*.sh
	@./scripts/init.sh
PHONY: init

lint:
	@./scripts/lint.sh
PHONY: lint

test: lint
	@./scripts/test.sh
PHONY: test

build: test
	@./scripts/build.sh
PHONY: build

run: test
	@./scripts/run.sh
PHONY: run
