#include "../libft.h"
#define maxlen(a, b) (a <= b) ? a : b;

char *ft_strsub(const char *str, unsigned int start, size_t len)
{
    size_t i_len = my_strlen((char *)str);
    size_t max_len = maxlen(len, i_len - start);

    if (start >= i_len)
    {
        max_len = 0;
    }
    char *s_tmp;
    if (!(s_tmp = (char *)malloc(max_len + 1)))
    {
        return NULL;
    }
    size_t j = 0;
    for (; j < max_len; j++)
    {
        s_tmp[j] = str[start + j];
    }
    s_tmp[j] = '\0';
    return s_tmp;
}