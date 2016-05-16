
CC=gcc
CXX=g++
AR=ar
ARFLAGS=cru
CFLAGS= -g -c 
TARGET = libprint.a
SRCS := $(wildcard *.c)
OBJS := $(patsubst %.c, %.o, $(SRCS)) 

all : $(TARGET)

$(TARGET) : $(OBJS)
	$(AR) $(ARFLAGS) $@ $(OBJS) 
	ranlib $@

win:
	gcc -g -c -fPIC myprint.c -o myprint.o
	ar cru libprint.a  myprint.o  
	ranlib libprint.a


%.o : %.c
	$(CC) $(CFLAGS) $< -o $@

clean :
	rm -rf ./${TARGET} $(OBJS)
