all: opencl_kernels.go

opencl_kernels.go:
	go-bindata -o opencl_kernels.go -pkg ethash cl/kernel1.cl cl/kernel2.cl cl/kernel3.cl

clean:
	rm -rf opencl_kernels.go
