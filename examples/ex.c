#include <stdio.h>

int * getOne() {
   int i = 1;
   return &i;
}

int main() {
    int *a = getOne();   // 1

    printf("%d\n", *a);
    return 0;
}