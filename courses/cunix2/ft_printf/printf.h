#include <stdlib.h>
#include <stdarg.h>
#include <unistd.h>

void ft_printf(const char *format, ...);

unsigned long ft_strlen(const char *s);

int ft_atoi(const char *nptr);
char *ft_itoa(int nmb);

char *s_float(const char *s, unsigned long width, int flags[4]);
char *s_conversion(const char *s, unsigned long width, int flags[4]);
char *i_conversion(int i, unsigned long width, int flags[4]);
