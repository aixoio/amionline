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
final: clean prod-build

clean:
	$(GOC) clean

.PHONY: dev all prod run-dev run-prod build prod-build final
.DEFAULT_GOAL = all
