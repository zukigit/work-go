#include "../headers/test.h"
#include "../headers/common.h"

extern int G1;
int add(int a, int b) {
    return a + b + G1;
}