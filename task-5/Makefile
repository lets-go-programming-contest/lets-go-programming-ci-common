gotest:
	@echo "Generate test suite..."
	go -C "${STATIC_DIR}" generate ./...

	@echo "Copy test suite from ${STATIC_DIR}/tests into ${SOURCE_DIR}/tests/ci"
	mkdir -p ${SOURCE_DIR}/tests/ci
	cp ${STATIC_DIR}/tests/task5_test.go ${SOURCE_DIR}/tests/ci/task5_test.go

	@echo "Run tests"
	go -C ${SOURCE_DIR} mod tidy
	go -C ${SOURCE_DIR} test -v --cover ./...
