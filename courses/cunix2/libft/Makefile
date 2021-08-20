CC					:= gcc
RM					:= rm -rf
CFLAGS		    	+= -std=c99 -Iinclude -Wall -Wextra -Werror -c
SRCS				:= $(wildcard src/*.c)
MY_SRCS				:= $(wildcard my_src/*.c)
OBJS				:= $(SRCS:.c=.o)
MY_OBJS				:= $(MY_SRCS:.c=.o)
LIB_NAME			:= libft.a
STYLE50				:= /usr/local/bin/style50

all:		    clean $(OBJS) $(MY_OBJS)
				@ar rcs $(LIB_NAME) src/*.o my_src/*.o

checkstyle:
				@$(STYLE50) $(SRCS) $(MY_SRCS) #&>/dev/null && echo "LOOKS GOOD!!!"

clean: clean_obj clean_lib

clean_obj:
				$(RM) $(OBJS) $(MY_OBJS)

clean_lib:
				$(RM) $(LIB_NAME)
