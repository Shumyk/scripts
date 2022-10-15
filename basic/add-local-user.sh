#!/bin/bash
#
# creates new user on local system for provided parameters

# verify superuser privileges
if [[ $(id -u) -ne 0 ]]
then
	echo "you should run this as super user" >&2
	exit 1
fi

if [[ ${#} -lt 1 ]]
then
	echo "account name is required" >&2
	echo "usage: ${0} <account name> [account info]..." >&2
	exit 1
fi


USERNAME=${1}
shift
ACCOUNT_INFO=${*}
PASSWORD=$(date +%s%N | sha256sum | head -c48)

# create user
adduser --create-home --comment "${ACCOUNT_INFO}" ${USERNAME} &> /dev/null
if [[ ${?} -ne 0 ]]
then
	echo "could not create user [${USERNAME}] with account info [${ACCOUNT_INFO}]" >&2
	exit 1
fi

# set password
echo ${PASSWORD} | passwd --stdin ${USERNAME} &> /dev/null
if [[ ${?} -ne 0 ]]
then
	echo "could not set password [${PASSWORD}] for user [${USERNAME}]" >&2
	exit 1
fi

# expire password on first login
passwd -e ${USERNAME} &> /dev/null
if [[ ${?} -ne 0 ]]
then
	echo "could not expire password for user [${USERNAME}]" >&2
	exit 1
fi

# print user info
echo "user name:"
echo ${USERNAME}
echo
echo "password:"
echo ${PASSWORD}
echo
echo "host:"
echo ${HOSTNAME}
