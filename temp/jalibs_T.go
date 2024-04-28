package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: ./src/jalibs/libjacommon_a-jalockutil.o ./src/libs/zbxcommon/str.o ./src/jalibs/libjacommon_a-jafile.o ./src/libs/zbxlog/log.o ./tools/json-c-0.9/json_object.o ./tools/json-c-0.9/printbuf.o ./tools/json-c-0.9/arraylist.o ./tools/json-c-0.9/linkhash.o ./src/libs/zbxsys/threads.o ./src/libs/zbxcommon/misc.o ./src/libs/zbxsys/mutexs.o
#include "./jainclude/jalockutil.h"
char	*CONFIG_LOG_FILE	 = "/var/log/jobarranger/jobarg_server.log";
int CONFIG_DB_CON_COUNT = 10;
char	*CONFIG_TMPDIR = NULL;
int	CONFIG_LOG_FILE_SIZE	= 1;
char	*CONFIG_FILE		 = "/home/zuki/Documents/dev/work/jobarranger-6.0.5.1/jaconf/jobarg_server.conf";

const char	title_message[] = "Job Arranger Server";
const char	*progname = NULL;
const char	usage_message[] = "[-hV] [-c <file>]";
const char	*help_message[] = {
	"Options:",
	"  -c --config <file>   Absolute path to the configuration file",
	"  -f --foreground     Run Job Arranger server in foreground",
	"",
	"Other options:",
	"  -h --help            Give this help",
	"  -V --version         Display version number",
	NULL
};
*/
import "C"

func main() {
	var folder_path *C.char = nil
	C.get_jaz_folder_path(folder_path)
}
