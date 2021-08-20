#include "../libft.h"

void *ft_memmove(void *dest, const void *src, size_t n)
{
    char *tmp_dest = dest;
    char tmp_src[n];
    ft_memcpy(tmp_src, src, n);
    for (size_t i = 0; i < n; i++)
    {
        tmp_dest[i] = tmp_src[i];
    }
    return dest;
}