.PHONY: atlas-init atlas-migrate atlas-apply atlas-status atlas-visualize atlas-help

DB_URL=postgres://postgres:password@localhost:5432/users_database?sslmode=disable

ENV_NAME=gorm

CMD := $(firstword $(MAKECMDGOALS))
ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))

define validate_args
  if [ -z "$(ARGS)" ]; then \
	  echo "✗ missing args, usage: make $(CMD) <ARGS>"; \
	  exit 1; \
	fi;
endef

# Generate new migration file
atlas-migrate:
	atlas migrate diff --env $(ENV_NAME)

atlas-apply:
	atlas migrate apply --env $(ENV_NAME) --url $(DB_URL)

atlas-apply-version:
	@$(validate_args)
	atlas migrate down --env $(ENV_NAME) --url $(DB_URL) --to-version $(ARGS)
atlas-status:
	atlas migrate status --env $(ENV_NAME) --url $(DB_URL)


#atlas-visualize:
#	atlas schema inspect --env $(ENV_NAME) --visualize --url $(DB_URL)


atlas-init:
	atlas schema inspect --env $(ENV_NAME)

# Help
atlas-help:
	@echo "Atlas Migration Make Commands:"
	@echo "  make atlas-init            - Initialize Atlas schema"
	@echo "  make atlas-migrate name=X  - Create a new migration named X"
	@echo "  make atlas-apply           - Apply pending migrations"
	@echo "  make atlas-status          - Show migration status"
	@echo "  make atlas-visualize       - Generate visual representation of schema"