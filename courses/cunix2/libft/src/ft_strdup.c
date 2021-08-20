#include "../libft.h"

char *ft_strdup(const char *s)
{
    char *tmp = (char *)malloc(my_strlen((char *)s) + 1);

    int i;
    for (i = 0; s[i]; i++)
    {
        tmp[i] = s[i];
    }
    tmp[i] = '\0';

    return tmp;
}
