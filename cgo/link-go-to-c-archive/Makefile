
libhello.a:
	go build -buildmode=c-archive -o libhello.a hello.go

libhello.so:
	go build -buildmode=c-shared -o libhello.so hello.go

clean:
	rm -rf libhello.*
