default: bin/go-workshop

bin/go-workshop: ./cmd/go-workshop/*.go
	mkdir -p bin
	go build -o $@ $^

# Create Realize config from template
.realize.yaml: etc/realize.yaml
	mkdir -p $(dir $@) && sed -e "s#%PROJECT_ROOT%#$(shell pwd)#g" $< > $@

# Requires Realize >= 2.0
# go get -u github.com/tockins/realize
dev: .realize.yaml
	realize start

deps:
	dep ensure -v

godeps: deps

test:
	go test ./cmd/... ./pkg/...

clean:
	rm -rf bin
