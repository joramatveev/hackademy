#include "../printf.h"

char *s_float(const char *s, unsigned long width, int flags[4])
{
    unsigned long str_len = ft_strlen(s);

    if (width <= str_len)
    {
        width = str_len;
    }

    char *new_s = (char *)malloc(sizeof(char) * (width + 1));
    new_s[width] = '\0';
    char fill_chr = ' ';

    if (flags[0])
    {
        fill_chr = '0';
    }

    if (flags[1])
    {
        for (unsigned long i = str_len; i < width; i++)
        {
            new_s[i] = fill_chr;
        }

        for (unsigned long i = 0; i < str_len; i++)
        {
            new_s[i] = s[i];
        }
    }
    else
    {
        unsigned long start = width - str_len;

        for (unsigned long i = 0; i < start; i++)
        {
            new_s[i] = fill_chr;
        }

        for (unsigned long i = start; i < width; i++)
        {
            new_s[i] = s[i - start];
        }
    }

    return new_s;
}
