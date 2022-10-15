#!/bin/bash
#
# this script deletes a user

# verifies superuser privileges
if [[ "${UID}" -ne 0 ]]
then
	echo 'please run with sudo or as root' >&2
	exit 1
fi

# assume the first argument is the user to delete
USER="${1}"

# delete the user
userdel ${USER}

# make sure the user got deleted
if [[ "${?}" -ne 0 ]]
then
	echo "The account ${USER} was NOT deleted" >&2
	exit 1
fi

# tell the user the account was deleted
echo "The account ${USER} was deleted"

exit 0
