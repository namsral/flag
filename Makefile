.PHONY: update

all: update

update:
	curl -LO https://raw.githubusercontent.com/golang/go/master/src/flag/example_test.go
	curl -LO https://raw.githubusercontent.com/golang/go/master/src/flag/example_value_test.go
	curl -LO https://raw.githubusercontent.com/golang/go/master/src/flag/export_test.go
	curl -LO https://raw.githubusercontent.com/golang/go/master/src/flag/flag.go
	curl -LO https://raw.githubusercontent.com/golang/go/master/src/flag/flag_test.go
