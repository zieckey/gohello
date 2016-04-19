// file hello.c
#include <stdio.h>
#include "libhello.h"

int main() {
  printf("This is a C Application.\n");
  GoString name = {(char*)"Jane", 4};
  SayHello(name);
  GoSlice buf = {(void*)"Jane", 4, 4};
  SayHelloByte(buf);
  SayBye();
  return 0;
}
