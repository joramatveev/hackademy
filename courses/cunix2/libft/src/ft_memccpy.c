#include "../libft.h"

void *ft_memccpy(void *dest, const void *src, int c, size_t n)
{
    char *tmp_dest = dest;
    const char *tmp_src = src;

    for (size_t i = 0; i < n; i++)
    {
        tmp_dest[i] = tmp_src[i];
        if (tmp_src[i] == (char)c)
        {
            return tmp_dest + i + 1;
        }
    }
    return NULL;
}