#!/bin/bash
#
# executes provided arguments as commands on all servers listed in default/supplied file

SSH_OPTIONS='-o ConnectTimeout=2'       # Options for the ssh command.
SERVER_FILE='/vagrant/servers'          # A list of servers, one per line.
DRY_RUN='false'
SUDO=''
VERBOSE='false'

# Display the usage and exit.
usage() {
	echo "Usage:" >&2
	echo "	${0} [OPTION]... COMMAND" >&2
	echo 'Description:' >&2
	echo '	executes all arguments as a single command on every server listed in the /vagrant/servers file by default' >&2
	echo >&2
	echo '	-f file' >&2
	echo '		overrides the default servers file of /vagrant/servers' >&2
	echo '	-n' >&2
	echo '		dry run, the commands will be displayed instead of executed' >&2
	echo '	-s' >&2
	echo '		runs the commands with superuser privileges on the remote servers' >&2
	echo '	-v' >&2
	echo '		verbose mode. displays the name of the server for which the command is being executed on' >&2
	exit 1
}

# Make sure the script is not being executed with superuser privileges.
if [[ "${UID}" -eq 0 ]]
then
	echo 'should not be executed with superuser privileges.' >&2
	echo 'specify -s option to execute remote commands with superuser privileges.' >&2
	usage
fi

# Parse the options.
while getopts f:nsv OPTION
do
	case ${OPTION} in
		f) SERVER_FILE="${OPTARG}" ;;
		n) DRY_RUN='true' ;;
		s) SUDO='sudo' ;;
		v) VERBOSE='true' ;;
		?) usage ;;
	esac
done

# Remove the options while leaving the remaining arguments.
shift "$(( OPTIND - 1 ))"
# If the user doesn't supply at least one argument, give them help.
if [[ "${#}" -eq 0 ]]
then
	usage
fi

# Anything that remains on the command line is to be treated as a single command.
COMMANDS="${*}"

# Make sure the SERVER_FILE file exists.
if [[ ! -e "${SERVER_FILE}" ]]
then
	echo "cannot open server list file ${SERVER_FILE}." >&2
	exit 1
fi

# Expect the best, prepare for the worst.
EXIT_STATUS=0

# Loop through the SERVER_LIST
for SERVER in $(cat ${SERVER_FILE})
do
	if [[ "${VERBOSE}" = 'true' ]]
	then
		echo "${SERVER}"
	fi

	SSH_COMMAND="ssh ${SSH_OPTIONS} ${SERVER} ${SUDO} ${COMMANDS}"

	# If it's a dry run, don't execute anything, just echo it.
	if [[ "${DRY_RUN}" = 'true' ]]
	then
		echo "DRY RUN: ${SSH_COMMAND}"
	else
		${SSH_COMMAND}
		SSH_EXIT_STATUS=${?}

    # Capture any non-zero exit status from the SSH_COMMAND and report to the user.
		if [[ "${SSH_EXIT_STATUS}" -ne 0 ]]
		then
		  EXIT_STATUS=${SSH_EXIT_STATUS}
			echo "commands [${COMMANDS}] was not able to run successfully on [${SERVER}]" >&2
		fi
	fi
done

exit ${EXIT_STATUS}
