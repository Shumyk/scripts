#!/bin/bash
#
# This script disables, deletes, and/or archives users on the local system.

MODE='disable'
REMOVE_HOME_DIRECTORY=''
ARCHIVE_HOME_DIRECTORY='false'

# Display the usage and exit.
usage() {
	echo "Usage: ${0} [-d] [-r] [-a] USER [USERN]..." >&2
	echo "Disables (expires/locks) accounts by default" >&2
	echo "Possible options:" >&2
	echo "	-d	deletes accounts instead of disabling them" >&2
	echo "	-r	removes home directory associated with the accounts" >&2
	echo "	-a	create an archive of the home directory associated with the accounts" >&2
	exit 1
}

# Make sure the UID of the account is at least 1000.
validateDeleteOperation() {
	local USERNAME=${1}
	local ACCOUNT_ID=$(id -u ${USERNAME})
	if [[ "${ACCOUNT_ID}" -lt "1000" ]]
	then
		echo "refuse to delete account [${USERNAME}], ID: [${ACCOUNT_ID}]" >&2
		echo "system accounts should be modified by system administrators" >&2
		exit 1
	fi
}

# Create an archive if requested to do so.
archiveHomeDirectory() {
	local USERNAME="${1}"

	if [[ "${ARCHIVE_HOME_DIRECTORY}" = 'true' ]]
	then
		local ARCHIVE_DIR='/archive/'
		local ARCHIVE_NAME="${ARCHIVE_DIR}${USERNAME}.$(date +%F.%N).tar.zip"

		mkdir -p "${ARCHIVE_DIR}"
		tar -zcf "${ARCHIVE_NAME}" "/home/${USERNAME}"

		echo "home directory for ${USERNAME} was archived to ${ARCHIVE_NAME}"
	fi
}

processUserOperation() {
	local USERNAME="${1}"

  validateDeleteOperation "${USERNAME}"
  if [[ "${MODE}" = 'disable' ]]
  then
    chage -E 0 ${USERNAME}
		echo "account ${USERNAME} was disabled"
  else
		archiveHomeDirectory "${USERNAME}"
    userdel ${REMOVE_HOME_DIRECTORY} "${USERNAME}"
		echo "account ${USERNAME} was deleted"
		if [[ -n "${REMOVE_HOME_DIRECTORY}" ]]; then echo "home directory removed for account"; fi
  fi
}

# Make sure the script is being executed with superuser privileges.
if [[ "${UID}" -ne "0" ]]
then
	echo 'you should run this as root' >&2
	exit 1
fi

# Parse the options.
while getopts dra OPTION
do
	case ${OPTION} in
		d) MODE='delete' ;;
		r) REMOVE_HOME_DIRECTORY='-r' ;;
		a) ARCHIVE_HOME_DIRECTORY='true' ;;
		?) usage ;;
	esac
done

# Remove the options while leaving the remaining arguments.
shift "$(( OPTIND - 1 ))"
# If the user doesn't supply at least one argument, give them help.
if [[ "${#}" -eq "0" ]]
then
	usage
fi

# Loop through all the usernames supplied as arguments.
while [[ "${#}" -gt "0" ]]
do
  echo "Processing user: ${USERNAME}"
	processUserOperation "${1}"
	shift
done

exit 0
