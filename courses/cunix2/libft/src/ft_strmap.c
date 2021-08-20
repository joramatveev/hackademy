#include "../libft.h"

char *ft_strmap(const char *str, char (*fn)(char))
{
    char *tmp = (char *)malloc(my_strlen((char *)str) + 1);
    int i;
    for (i = 0; str[i]; i++)
    {
        tmp[i] = (*fn)(str[i]);
    }
    tmp[i] = '\0';

    return tmp;
}