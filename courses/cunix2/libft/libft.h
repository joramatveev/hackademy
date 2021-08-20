#include <unistd.h>
#include <stdlib.h>

void ft_bzero(void *s, size_t n);
int ft_isalpha(int chr);
int ft_isascii(int chr);
int ft_isdigit(int chr);
char *ft_strchr(const char *s, int chr);
char *ft_strdup(const char *s);
int ft_strncmp(const char *s1, const char *s2, size_t n);
char *ft_strrchr(const char *s, int c);
int ft_tolower(int chr);
int ft_toupper(int chr);

int ft_abs(int i);
div_t ft_div(int numer, int denom);
char* ft_strstr(const char* s1, const char* s2);
char *ft_strnstr(const char *haystack, const char *needle, size_t len);
void *ft_memset(void *s, int c, size_t n);
void *ft_memcpy(void *dest, const void *src, size_t n);
void *ft_memccpy(void *dest, const void *src, int c, size_t n);
void *ft_memmove(void *dest, const void *src, size_t n);
void *ft_memchr(const void *s, int c, size_t n);
int ft_memcmp(const void *s1, const void *s2, size_t n);

void ft_striter(char *str, void (*fn)(char *));
char *ft_strmap(char const *str, char (*fn)(char));
char *ft_strsub(char const *str, unsigned int start, size_t len);
char *ft_strjoin(char const *s1, char const *s2);
char *ft_strtrim(char const *s);
char **ft_strsplit(char const *s, char c);

// MY SRC'S
unsigned int my_strlen(char const *str);
int my_strcmp(const char *s1, const char *s2);
char *my_strcpy(char *dest, const char *src);
int my_atoi(const char *nptr);
char* my_itoa(int nmb);
int my_puts(const char *s);