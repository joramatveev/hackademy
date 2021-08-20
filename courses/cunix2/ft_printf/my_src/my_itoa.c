#include <stdlib.h>

char *my_itoa(int num)
{
    char buff[32];
    int i = 0, ok = 0;
    if (num < 0)
    {
        num *= -1;
        ok = 1;
    }
    do
    {
        buff[i++] = num % 10 + '0';
        num /= 10;
    }
    while (num != 0);

    char *res = malloc(i + ok + 1);
    if (ok)
    {
        res[0] = '-';
    }
    for (int j = i - 1; j >= 0; j--)
    {
        res[i - j - 1 + ok] = buff[j];
    }
    res[i + ok] = '\0';
    return res;
}