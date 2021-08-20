#include "../libft.h"

int ft_strncmp(const char *s1, const char *s2, size_t n)
{
    if (n == 0)
    {
        return 0;
    }
    for (size_t i = 1; i <= n && *s1++ == *s2++; i++);
    return (*--s1 < *--s2) ? -1 : (*s1 > *s2);
}