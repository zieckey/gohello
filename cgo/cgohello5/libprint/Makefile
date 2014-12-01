
CC=gcc
CXX=g++
AR=ar
ARFLAGS=cru
CFLAGS= -g -c -fPIC
TARGET = libprint.a
SRCS := $(wildcard *.c)
OBJS := $(patsubst %.c, %.o, $(SRCS)) 

all : $(TARGET)
	go build
	./${TARGET}

$(TARGET) : $(OBJS)
	$(AR) $(ARFLAGS) $@ $(OBJS) 
	ranlib $@

%.o : %.c
	$(CC) $(CFLAGS) $< -o $@

clean :
	rm -rf ./${TARGET}
