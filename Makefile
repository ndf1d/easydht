LINT_PATH := ./...
fetch:
	govendor sync -v

lint_all:
	gometalinter --vendor ./...

lint_path:
	gometalinter --vendor ${LINT_PATH}
