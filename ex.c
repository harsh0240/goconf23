#include <stdio.h>

int * getRandom( ) {
   int i = 5;
   return &i;
}

int anotherRandom() {
    int i = 6;
    return i;
}

int main () {
    int *a = getRandom();       // 5
    int b = anotherRandom();    // 6

    printf("%d %d\n", *a, b);
    return 0;
}