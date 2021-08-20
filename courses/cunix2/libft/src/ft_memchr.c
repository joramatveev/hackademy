#include "../libft.h"

void *ft_memchr(const void *s, int c, size_t n)
{
    const char *tmp = s;
    for (size_t i = 0; i < n; i++)
    {
        if (tmp[i] == (char)c)
        {
            return (void *)(tmp + i);
        }
    }
    return NULL;
}