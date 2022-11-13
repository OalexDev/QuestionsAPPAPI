build:
	@echo "╰─❯  Build Application"
	@sam build

run: build
	@echo "╰─❯  Running Application"
	@sam local start-api --env-vars env.json
