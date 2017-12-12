y.go: smiles.y
	go tool yacc -p smiles smiles.y

fmt:
	gofmt -s=true -tabs=false -tabwidth=2 -w=true `find . -name "*.go"`
