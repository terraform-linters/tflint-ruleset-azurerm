default: build

test:
	go test $$(go list ./... | grep -v integration)

e2e:
	cd integration && go test && cd ../

build:
	go build

install: build
	mkdir -p ~/.tflint.d/plugins
	mv ./tflint-ruleset-azurerm ~/.tflint.d/plugins

lint:
	golint --set_exit_status $$(go list ./...)
	go vet ./...

tools:
	go install golang.org/x/lint/golint

.PHONY: test e2e build install lint tools
