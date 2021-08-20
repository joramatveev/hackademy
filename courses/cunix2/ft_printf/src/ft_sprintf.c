#include "../printf.h"

int ft_sprintf(char *arr, const char *format, ...)
{
    va_list arg;
    va_start(arg, format);

    unsigned int written_chars = 0;
    unsigned long str_len = my_strlen(format);
    int flags[4] = {0, 0, 0, 0 }; // [0] = 0, [1] = -, [2] = +, [3] = space
    unsigned long width = 0;

    for (unsigned long i = 0; i < str_len; i++)
    {
        while (format[i] != '%' && i < str_len)
        {
            arr[written_chars] = format[i];
            written_chars++;
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

            // width
            if (format[i] >= '1' && format[i] <= '9')
            {
                width = my_atoi(&format[i]);
            }

            while (format[i] >= '0' && format[i] <= '9')
            {
                i++;
            }

            // type
            char *inserted_data = NULL;

            switch (format[i])
            {
                case 'i':
                case 'd':
                    inserted_data = i_conversions(va_arg(arg, int), width, flags);
                    break;
                case 'c':
                    inserted_data = c_conversions(va_arg(arg, int), width, flags);
                    break;
                case 's':
                    inserted_data = s_conversions(va_arg(arg, char *), width, flags);
                    break;
                case '%':
                    inserted_data = "%";
                    break;
                default:
                    exit(1);
            }

            unsigned int inserted_len = my_strlen(inserted_data);

            for (unsigned int i = 0; i < inserted_len; i++)
            {
                arr[written_chars] = inserted_data[i];
                written_chars++;
            }
        }
    }
    arr[written_chars++] = '\0';

    return written_chars;
}
