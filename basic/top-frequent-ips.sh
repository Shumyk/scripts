#!/bin/bash
#
# forms a list of IPs of top frequent callers

LIMIT=10
LOG_FILE=${1}


if [[ ! -r "${LOG_FILE}" ]]
then
	echo "provided file ${LOG_FILE} does not exists or not readable" >&2
	exit 1
fi

echo 'Count,IP,Location'
grep "Failed" ${LOG_FILE} | awk '{print $(NF - 3)}' | sort | uniq -c | sort -n | while read COUNT IP
do
	if [[ "${COUNT}" -gt "${LIMIT}" ]]
	then
		LOCATION=$(geoiplookup "${IP}" | awk -F ', ' '{print $2}')
		echo "${COUNT},${IP},${LOCATION}"
	fi
done

exit 0
