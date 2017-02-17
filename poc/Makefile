.PHONY: init build-base test

init:
	@if [ ! -d "gcp-php-docker" ]; then git clone https://github.com/GoogleCloudPlatform/php-docker.git gcp-php-docker; fi

build-base:
	docker build -t engine/php gcp-php-docker/php-nginx

test:
	./e2e_test.sh
