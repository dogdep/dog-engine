.PHONY: dogapi dog

dogapi:
	@cd cmd/dogapi && go build -ldflags="-s -w" -o ../../bin/dogapi

dog:
	@cd cmd/dog && go build -ldflags="-s -w" -o ../../bin/dog
