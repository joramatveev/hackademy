int my_atoi(const char *nptr)
{
    int res = 0, sign = 1;
    if (*nptr == '-')
    {
        sign = -1;
        nptr++;
    }
    while (*nptr != '\0')
    {
        if (*nptr < '0' || *nptr > '9')
        {
            break;
        }
        res = res * 10 + (*nptr - '0');
        nptr++;
    }
    return res * sign;
}