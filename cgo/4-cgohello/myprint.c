#include "myprint.h"
#include <stdio.h>

extern FunctionExportedFromGo();

void myprint(const char* str)
{
  printf("%s\n", str);
  FunctionExportedFromGo();
  GoFuncPrintxx();
}
