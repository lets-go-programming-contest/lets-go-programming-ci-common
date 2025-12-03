check-coverage:
	@echo "Copy source code from ${STATIC_DIR}/tests/internal into ${SOURCE_DIR}/internal"

	cp ${STATIC_DIR}/tests/internal/db/db.go ${SOURCE_DIR}/internal/db/db.go
	cp ${STATIC_DIR}/tests/internal/wifi/wifi.go ${SOURCE_DIR}/internal/wifi/wifi.go

	@echo "Check coverage"
	@coverage_db=$$(go -C ${SOURCE_DIR} test --cover ./internal/db/... | grep -oP 'coverage:\s+\K\d+(\.\d+)?%'); \
	if [ "$$coverage_db" != "100.0%" ]; then \
		echo "ERROR: Coverage for internal/db is $$coverage_db, expected 100%"; \
		exit 1; \
	fi && \
	@coverage_wifi=$$(go -C ${SOURCE_DIR} test --cover ./internal/wifi/... | grep -oP 'coverage:\s+\K\d+(\.\d+)?%'); \
	if [ "$$coverage_wifi" != "100.0%" ]; then \
		echo "ERROR: Coverage for internal/db is $$coverage_wifi, expected 100%"; \
		exit 1; \
	fi