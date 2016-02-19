
CC := gcc
CFLAGS := -g -c -fPIC -Wshadow -Wcast-qual -Wcast-align -Wwrite-strings -Wsign-compare -Winvalid-pch -fms-extensions -Wall

hello:	libhello.a hello.o
	$(CC) hello.o -L. ./libhello.a -lpthread -o $@

libhello.a: hello.go
	go build -buildmode=c-archive -o libhello.a hello.go

libhello.so:
	go build -buildmode=c-shared -o libhello.so hello.go

%.o : %.c
	$(CC) $(CFLAGS) $< -o $@

clean:
	rm -rf libhello.* *.o hello
