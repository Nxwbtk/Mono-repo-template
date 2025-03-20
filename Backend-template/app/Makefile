
new:
	@if [ -z "$(word 2, $(MAKECMDGOALS))" ]; then \
		echo "Error: You must specify a module name. Use 'make new <name>'"; \
		exit 1; \
	fi
	@MODULE_NAME=$(word 2, $(MAKECMDGOALS)); \
	LOWER_NAME=$$(echo $$MODULE_NAME | tr '[:upper:]' '[:lower:]'); \
	PASCAL_NAME=$$(echo $$MODULE_NAME | sed -e 's/\b\(.\)/\u\1/g'); \
	echo "Creating new module structure for $$PASCAL_NAME"; \
	cd services; \
	mkdir -p $$MODULE_NAME; \
	sleep 1; \
	touch $$MODULE_NAME/$$LOWER_NAME.handler.go; \
	touch $$MODULE_NAME/$$LOWER_NAME.repo.go; \
	touch $$MODULE_NAME/$$LOWER_NAME.service.go; \
	echo "package $$LOWER_NAME\n\ntype $${PASCAL_NAME}Handler struct {\n\t$${PASCAL_NAME}Service $${PASCAL_NAME}Service\n}\n\nfunc New$${PASCAL_NAME}Handler($${LOWER_NAME}Service $${PASCAL_NAME}Service) *$${PASCAL_NAME}Handler {\n\treturn &$${PASCAL_NAME}Handler{\n\t\t$${PASCAL_NAME}Service: $${LOWER_NAME}Service,\n\t}\n}" > $$MODULE_NAME/$$LOWER_NAME.handler.go; \
	echo "package $$LOWER_NAME\n\nimport (\n\t\"gorm.io/gorm\"\n)\n\ntype $${PASCAL_NAME}Repo interface {\n}\n\ntype $${LOWER_NAME}Repo struct {\n\tdb *gorm.DB\n}\n\nfunc New$${PASCAL_NAME}Repo(db *gorm.DB) $${PASCAL_NAME}Repo {\n\treturn &$${LOWER_NAME}Repo{\n\t\tdb: db,\n\t}\n}" > $$MODULE_NAME/$$LOWER_NAME.repo.go; \
	echo "package $$LOWER_NAME\n\ntype $${PASCAL_NAME}Service interface {\n}\n\ntype $${LOWER_NAME}Service struct {\n\t$${LOWER_NAME}Repository $${PASCAL_NAME}Repo\n}\n\nfunc New$${PASCAL_NAME}Service($${LOWER_NAME}Repository $${PASCAL_NAME}Repo) $${PASCAL_NAME}Service {\n\treturn &$${LOWER_NAME}Service{\n\t\t$${LOWER_NAME}Repository: $${LOWER_NAME}Repository,\n\t}\n}" > $$MODULE_NAME/$$LOWER_NAME.service.go; \
	echo "\n\033[1;32m┌────────────────────────────────────────────┐"; \
	echo "│ █▄─▀█▀─▄█▄─▄█─▄▄▄▄█─▄▄▄▄█▄─▄█─▄▄─█▄─▀█▄─▄█ │"; \
	echo "│ ██─█▄█─███─██▄▄▄▄─█▄▄▄▄─██─██─██─██─█▄▀─██ │"; \
	echo "│ █▄▄▄█▄▄▄█▄▄▄█▄▄▄▄▄█▄▄▄▄▄█▄▄▄█▄▄▄▄█▄▄▄██▄▄█ │"; \
	echo "│                 COMPLETED                  │"; \
	echo "│ Module $$PASCAL_NAME executed successfully! │"; \
	echo "│ Good hunting, Stay frosty! 🚀              │"; \
	echo "└────────────────────────────────────────────┘\033[0m\n";
%:
	@:
dev:
	@echo "Running dev server"
	@nodemon --exec go run $(word 2, $(MAKECMDGOALS)) --signal SIGTERM

gen-docs:
	@echo "Generating docs"
	@swag init

reload: gen-docs dev $(word 2, $(MAKECMDGOALS))

d-test:
	rm -rf services/test/

.PHONY: new dev gen-docs d-test