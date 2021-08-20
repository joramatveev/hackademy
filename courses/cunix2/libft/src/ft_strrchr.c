#include "../libft.h"

char *ft_strrchr(const char *s, int c)
{
    c = (char)c;
    int j;
    for (j = my_strlen((char *)s); j >= 0; j--)
    {
        if (s[j] == (char)c)
        {
            return (char *)(s + j);
        }
    }
    return NULL;
}
