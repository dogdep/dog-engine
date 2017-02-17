# POC. Simulate Lumen php micro-framework deployment

Main goal is to simulate theoretical deployment process of an php app
as a single container from developers perspective.

UC. Developer has an amazing app called fluffy which is built using a
php framework and wants to easily deploy to staging environment with minimal
work to be done.

1. Developer types in command `./dog.sh deploy` in application source root
2. An archive is created and sent over to the _engine_
3. _Engine_ extracts it, builds the app container, starts the container
4. Developer can happily see the app running.

## Prerequisites

- docker is set up on your local machine
- make

## Walkthrough

Below are steps to get ready the fake _engine_.

Build the base docker image via `make init`, which clones repository
[GoogleCloudPlatform/php-docker](https://github.com/GoogleCloudPlatform/php-docker/).
Then executing `make build-base` builds locally the base image as `engine/php` image
(this will take a while to build).

Enter the directory `fluffy` and enter `./dog.sh deploy`. Afterwards
run shell script `./build_and_run_.sh deployments/fluffy.tar.gz fluffy`.

make `test` - initializes everything and runs an e2e test.

tl;dr:

3. `make test`
