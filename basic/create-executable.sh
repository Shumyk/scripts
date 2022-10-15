#!/bin/bash
#
# creates executable file

if [[ "${#}" -ne 1 ]]
then
	echo "Usage: ${0} <filename>" >&2
	echo '	only single filename is expected' >&2
	exit 1
fi

touch "${1}"
chmod 755 "${1}"

exit 0
