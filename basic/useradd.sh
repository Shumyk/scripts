#!/bin/bash
#
# this script creates an account on the local system
# you will be prompted for the account name and password

# ask for the username
read -p 'Enter the username to create: ' USERNAME

# ask for the real name
read -p 'Enter the name of the person who this account is for: ' COMMENT

# ask for the password
read -p 'Enter the password to use for the account: ' PASSWORD

# create the user
useradd -c "${COMMENT}" -m ${USERNAME}

# set the password for the user
echo ${PASSWORD} | passwd --stdin ${USERNAME}

# force password change on first login
passwd -e ${USERNAME}
