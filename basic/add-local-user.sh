#!/bin/bash
#
# This script creates a new user on the local system.
# You must supply a username as an argument to the script.
# Optionally, you can also provide a comment for the account as an argument.
# A password will be automatically generated for the account.
# The username, password, and host for the account will be displayed.

# Make sure the script is being executed with superuser privileges.
if [[ $(id -u) -ne 0 ]]
then
	echo "you should run this as super user" >&2
	exit 1
fi

# If the user doesn't supply at least one argument, then give them help.
if [[ ${#} -lt 1 ]]
then
	echo "account name is required" >&2
	echo "usage: ${0} <account name> [account info]..." >&2
	exit 1
fi

# The first parameter is the user name
USERNAME=${1}
# The rest of the parameters are for the account comments.
shift
ACCOUNT_INFO=${*}
# Generate a password.
PASSWORD=$(date +%s%N | sha256sum | head -c48)

# Create the user.
adduser --create-home --comment "${ACCOUNT_INFO}" ${USERNAME} &> /dev/null
# Check to see if the useradd command succeeded.
# We don't want to tell the user that an account was created when it hasn't been.
if [[ ${?} -ne 0 ]]
then
	echo "could not create user [${USERNAME}] with account info [${ACCOUNT_INFO}]" >&2
	exit 1
fi

# Set the password.
echo ${PASSWORD} | passwd --stdin ${USERNAME} &> /dev/null
# Check to see if the passwd command succeeded.
if [[ ${?} -ne 0 ]]
then
	echo "could not set password [${PASSWORD}] for user [${USERNAME}]" >&2
	exit 1
fi

# Force password change on first login.
passwd -e ${USERNAME} &> /dev/null
# check password expire operation success
if [[ ${?} -ne 0 ]]
then
	echo "could not expire password for user [${USERNAME}]" >&2
	exit 1
fi

# Display the username, password, and the host where the user was created.
echo "user name:"
echo ${USERNAME}
echo
echo "password:"
echo ${PASSWORD}
echo
echo "host:"
echo ${HOSTNAME}
