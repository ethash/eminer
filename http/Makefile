all: templates.go

templates.go:
	go-bindata -o templates.go -pkg http templates/

clean:
	rm -rf templates.go
