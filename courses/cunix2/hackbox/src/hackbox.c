#include <stdio.h>
#include <string.h>
#include <unistd.h>

#include "lib/cat.c"
#include "lib/cp.c"
#include "lib/ls.c"
#include "lib/mkdir.c"

static void
hackbox_welcome(void)
{
    eprintf("\nHackBox v0.0.1 multi-call binary.\n\n"
            "Try './hackbox --help' for more information.\n\n");
}

static void
hackbox_usage(void)
{
    eprintf("\nHackBox v0.0.1 multi-call binary.\n\n"
            "Usage: ./hackbox cat [-u] [file ...]\n"
            "   or: ./hackbox cp [-afpv] [-R [-H | -L | -P]\n"
            "   or: ./hackbox ls [-l] [path ...]\n"
            "   or: ./hackbox mkdir [-p] [-m mode] name ...\n"
            "   or: ./hackbox --help\n\n"
    );
}

int main(int argc, char **argv) {

    if (argc <= 1) {
        hackbox_welcome();
        return 0;
    }

    if (strcmp(argv[1], "cat") == 0 || strcmp(argv[1], "./cat") == 0) {
        argc--;
        argv++;
        cat_main(argc, argv);
    } else if (strcmp(argv[1], "cp") == 0 || strcmp(argv[1], "./cp") == 0) {
        argc--;
        argv++;
        cp_main(argc, argv);
    } else if (strcmp(argv[1], "ls") == 0 || strcmp(argv[1], "./ls") == 0) {
        argc--;
        argv++;
        ls_main(argc, argv);
    } else if (strcmp(argv[1], "mkdir") == 0 || strcmp(argv[1], "./mkdir") == 0) {
        argc--;
        argv++;
        mkdir_main(argc, argv);
    } else if (strcmp(argv[1], "help") == 0 || strcmp(argv[1], "./help") == 0 || strcmp(argv[1], "--help") == 0) {
        hackbox_usage();
    } else {
        hackbox_welcome();
    }

    return 0;
}