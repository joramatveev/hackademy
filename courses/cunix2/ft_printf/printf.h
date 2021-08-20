#include "../libft/libft.h"
#include <stdlib.h>
#include <stdarg.h>
#include <unistd.h>

void ft_printf(const char *format, ...);
char *s_float(const char *s, unsigned long width, int flags[4]);
char *c_conversions(char c, unsigned long width, int flags[4]);
char *s_conversions(const char *s, unsigned long width, int flags[4]);
char *i_conversions(int i, unsigned long width, int flags[4]);
int ft_sprintf(char *arr, const char *format, ...);