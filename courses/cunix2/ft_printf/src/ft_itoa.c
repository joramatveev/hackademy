#include "../printf.h"

char *ft_itoa(int nmb)
{
    const int int_char_offset = 48;
    long long bnmb = (long long)nmb;
    int str_len = 1, start = 0;

    if (bnmb < 0)
    {
        bnmb = -bnmb;
        start = 1;
        str_len++;
    }

    for (long long x = bnmb; x / 10 != 0; x /= 10)
    {
        str_len++;
    }

    char *str = (char *)malloc((str_len + 1) * sizeof(char));
    str[str_len] = '\0';

    if (start)
    {
        str[0] = '-';
    }

    for (int i = str_len - 1; i >= start; i--)
    {
        str[i] = bnmb % 10 + int_char_offset;
        bnmb /= 10;
    }

    return str;
}
