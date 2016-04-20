#ifndef PRINTS_HEAD
#define PRINTS_HEAD
void myprint(const char* str);

// the return pointer need to free(ed)
char* retmalloc(int len, int *retlen);
#endif
