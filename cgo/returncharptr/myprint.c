#include "myprint.h"
#include <stdio.h>
#include <stdlib.h>

extern FunctionExportedFromGo();

void myprint(const char* str)
{
  printf("%s\n", str);
  FunctionExportedFromGo();
  GoFuncPrintxx();
}


char* retmalloc(int len, int *retlen)
{
    char* p = malloc(len);
    *retlen = len;
    return p;
}
