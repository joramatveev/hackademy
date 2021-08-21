#include "../printf.h"

int ft_atoi(const char *nptr)
{
    const int int_to_char_offset = 48;
    int num = 0, minus = 1;

    if (*nptr == '-')
    {
        minus = -1;
        nptr++;
    }

    for (int i = 0; nptr[i] <= '9' && nptr[i] >= '0'; i++)
    {
        num *= 10;
        num += (nptr[i] - int_to_char_offset) * minus;
    }

    return num;
}
