CC			:= gcc
RM			:= rm -rf
CFLAGS		+= -std=c99 -Iinclude -Wall -Wextra -Werror -g -c
SRCS		:= $(wildcard src/*.c)
MY_SRCS		:= $(wildcard my_src/*.c)
OBJS		:= $(SRCS:.c=.o)
MY_OBJS		:= $(MY_SRCS:.c=.o)
LIB_NAME	:= printf.a
STYLE50		:= /usr/local/bin/style50


printf:		fclean $(OBJS) $(MY_OBJS)
			@$(CC) $(CFLAGS) $(SRCS) $(MY_SRCS)

all:		fclean $(OBJS) $(MY_OBJS)
			@$(CC) $(CFLAGS) $(SRCS) $(MY_SRCS)
			@ar rcs $(LIB_NAME) $(OBJS) $(MY_OBJS)

checkstyle:
			$(STYLE50) $(SRCS) $(MY_SRCS)

fclean:
			$(RM) $(OBJS) $(MY_OBJS) $(LIB_NAME)