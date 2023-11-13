GOC = go

GOCPROD = go
GOCPRODFLAGS = -ldflags "-s -w"

dev:
	$(GOC) build

prod:
	$(GOCPROD) build $(GOCPRODFLAGS)

run-dev:
	$(GOC) run .

run-prod:
	$(GOCPROD) run $(GOCPRODFLAGS) .

all: run-dev
build: dev
prod-build: prod

clean:
	$(GOC) clean

.PHONY: dev all prod run-dev run-prod build prod-build
.DEFAULT_GOAL = all
