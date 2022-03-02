all: cl.go gcn.go

cl.go:
	go-bindata -o cl -split -pkg cl cl/kernel1.cl

gcn.go:
	go-bindata -o gcn -split 	-pkg gcn gcn/bin/*

clean:
	rm -rf cl/*.go
	rm -rf gcn/*.go
