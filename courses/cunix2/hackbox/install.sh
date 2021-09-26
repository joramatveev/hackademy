#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'
{
	cd ./src && rm -f {hackbox,hackbox.o} && make && mv ./hackbox ../ &&
	echo -e "${GREEN}Installation completed successfully${NC}"
	
} || {
	echo -e "${RED}Installation completed with error${NC}"
}