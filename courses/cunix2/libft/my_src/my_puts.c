#include <unistd.h>

int my_puts(const char *s)
{
    char new_line = '\n';
    for (int i = 0; s[i] != '\0'; i++)
    {
        write(1, &s[i], sizeof(char));
    }
    write(1, &new_line, sizeof(char));

    return 1;
}