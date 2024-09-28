DOCKER_COMPOSE = docker-compose
GO = go
BINARY_NAME = pack-calculator
MAIN_PATH = ./cmd/main.go

.PHONY: test build run stop clean

test:
	$(GO) test -v ./...

build:
	$(GO) build -o $(BINARY_NAME) $(MAIN_PATH)

run:
	$(DOCKER_COMPOSE) up --build

stop:
	$(DOCKER_COMPOSE) down

deploy-gcp:
	docker build -t us-central1-docker.pkg.dev/pack-calculator-project/pack-calculator-repo/pack-calculator .
	docker push us-central1-docker.pkg.dev/pack-calculator-project/pack-calculator-repo/pack-calculator
	gcloud run deploy pack-calculator \
      --image us-central1-docker.pkg.dev/pack-calculator-project/pack-calculator-repo/pack-calculator \
      --platform managed \
      --region us-central1 \
      --allow-unauthenticated

clean:
	rm -f $(BINARY_NAME)
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans