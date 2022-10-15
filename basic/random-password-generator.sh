#!/bin/bash
#
# this script generates a random password
# this user can set the password length with -l and add a special character with -s
# verbose mode can be enabled with -v

VERBOSE='false'                  # set a default verbosity
LENGTH=48                        # set a default password length
USE_SPECIAL_CHARACTER='false'    # set a default for special character usage


usage() {
	echo "Usage: ${0} [-vs] [-l LENGTH]" >&2
	echo 'Generate a random password.'
	echo '	-l LENGTH	Specify the password length.'
	echo '	-s        Append a special character to the password.'
	echo '	-v		    Increase verbosity.'
	exit 1
}

log() {
	local MESSAGE="${*}"
	if [[ "${VERBOSE}" = 'true' ]]
	then
		echo "${MESSAGE}"
	fi
}


# parse options
while getopts vl:s OPTION
do
	case ${OPTION} in
		v)
			VERBOSE='true'
			log 'verbose mode on'
			;;
		l) LENGTH="${OPTARG}" ;;
		s) USE_SPECIAL_CHARACTER='true' ;;
		?) usage ;;
	esac
done

# remove the options while leaving the remaining arguments
shift "$(( OPTIND - 1 ))"
if [[ "${#}" -gt 0 ]]
then
	usage
fi

log 'Generating a password.'
PASSWORD=$(date +%s%N${RANDOM}${RANDOM} | sha256sum | head -c${LENGTH})

# append a special character if requested to do so.
if [[ "${USE_SPECIAL_CHARACTER}" = 'true' ]]
then
	log 'selecting a random special character'
	SPECIAL_CHARACTER=$(echo '!@#$%^&*()_+=-' | fold -w1 | shuf | head -n1)
	PASSWORD="${PASSWORD}${SPECIAL_CHARACTER}"
fi

# display the password
log 'Done.'
log 'Here is the password:'
echo "${PASSWORD}"

exit 0
