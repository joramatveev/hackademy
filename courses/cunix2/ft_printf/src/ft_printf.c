#include "../printf.h"

void print_buf(char *buf, int align, int len)
{
    int pad = 0;
    if (!buf)
    {
        len = 6;
    }
    pad = align - len;
    if (pad < 0)
    {
        pad = 0;
    }
    while (pad--)
    {
        write(1, " ", 1);
    }
    if (buf)
    {
        write(1, buf, len);
    }
    else
    {
        write(1, "(null)", 6);
    }
}

void print_char(char ch, int align)
{
    char *str_buf = (char *)malloc(2 * sizeof(char));
    str_buf[0] = ch;
    str_buf[1] = '\0';
    print_buf(str_buf, align, 1);
    free(str_buf);
}

void ft_printf(const char *format, ...)
{
    va_list arg;
    va_start(arg, format);

    unsigned long str_len = ft_strlen(format);
    int flags[4] = {0, 0, 0, 0};
    unsigned long width = 0;

    for (unsigned long i = 0; i < str_len; i++)
    {
        while (format[i] != '%' && i < str_len)
        {
            write(STDOUT_FILENO, &format[i], 1);
            i++;
        }

        if (format[i] == '%')
        {
            i++;
            while (format[i] == '0' || format[i] == '-' || format[i] == '+' || format[i] == ' ')
            {
                switch (format[i])
                {
                    case '0':
                        flags[0] = 1;
                        break;
                    case '-':
                        flags[1] = 1;
                        break;
                    case '+':
                        flags[2] = 1;
                        break;
                    default:
                        flags[3] = 1;
                }
                i++;
            }

            if (flags[1])
            {
                flags[0] = 0;
            }

            if (flags[2])
            {
                flags[3] = 0;
            }

            if (format[i] >= '1' && format[i] <= '9')
            {
                width = ft_atoi(&format[i]);
            }

            while (format[i] >= '0' && format[i] <= '9')
            {
                i++;
            }

            char *inserted_data = NULL;

            switch (format[i])
            {
                case 'i':
                case 'd':
                    inserted_data = i_conversion(va_arg(arg, int), width, flags);
                    break;
                case 'c':
                    print_char(va_arg(arg, int), width);
                    continue;
                    break;
                case 's':
                    inserted_data = s_conversion(va_arg(arg, char *), width, flags);
                    break;
                case '%':
                    inserted_data = "%";
                    break;
                default:
                    exit(1);
            }

            write(STDOUT_FILENO, inserted_data, ft_strlen(inserted_data));
        }
    }
}