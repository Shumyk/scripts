#!/bin/bash
#
# script backup requested file

# this function sends a message to syslog and to STDOUT if VERBOSE is true
log() {
	local MESSAGE="${*}"
	if [[ "${VERBOSE}" = 'true' ]]
	then
		echo "${MESSAGE}"
	fi
	logger -t ${0} "${MESSAGE}"
}

# this function creates a backup of a file. returns non-zero status on error.
backup_file() {
	local FILE="${1}"
	# make sure the file exists
	if [[ -f ${FILE} ]]
	then
		local BACKUP_FILE="/var/tmp/$(basename ${FILE}).$(date +%F-%N)"
		log "Backing up ${FILE} to ${BACKUP_FILE}"

		# the exit status of the function will be the exit status of the cp command
		cp -p ${FILE} ${BACKUP_FILE}
	else
		# the file does not exist, so return a non-zero exit status
		return 1
	fi
}


readonly VERBOSE='true'
backup_file "/etc/passwd"
# make a decision based on the exit status of the function
if [[ ${?} -eq 0 ]]
then
	log 'file backup succeeded!'
else
	log 'file backup failed!'
	exit 1
fi

