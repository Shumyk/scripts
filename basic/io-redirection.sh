#!/bin/bash
#
# this script demonstrates I/O redirection

# redirect STDOUT to a file
FILE="/tmp/data"
head -n1 /etc/passwd > ${FILE}

# redirect STDIN to a program
read LINE < ${FILE}
echo "LINE contains: ${LINE}"

# redirect STDOUT to a file, overwriting the file
head -n3 /etc/passwd > ${FILE}
echo
echo "Contents of ${FILE}:"
cat ${FILE}


# redirect STDOUT to a file, appending to the file
echo "${RANDOM} ${RANDOM}" >> ${FILE}
echo "${RANDOM} ${RANDOM}" >> ${FILE}
echo
echo "Contents of ${FILE}:"
cat ${FILE}

# redirect STDIN to a program, using FD 0
read LINE 0< ${FILE}
echo
echo "LINE contains: ${LINE}"

# redirect STDOUT to a file using FD 1, overwriting the file
head -n3 /etc/passwd 1> ${FILE}
echo
echo "Contents of ${FILE}:"
cat ${FILE}

# redirect STDERR to a file using FD 2
ERR_FILE="/tmp/data.err"
head -n3 /etc/passwd /fakefile 2> ${ERR_FILE}

# redirect STDOUT and STDERR to a file
head -n3 /etc/passwd /fakefile &> ${FILE}
echo
echo "New Contents of ${FILE}:"
cat ${FILE}


# redirect STDOUT and STDERR through a pipe
echo
head -n3 /etc/passwd /fakefile |& cat -n

# send output to STDERR
echo "This is ERROR!" >&2

# discard STDOUT
echo
echo "Discarding STDOUT:"
head -n3 /etc/passwd /fakefile > /dev/null

# discard STDERR
echo
echo "Discarding STDERR:"
head -n3 /etc/passwd /fakefile 2> /dev/null

# discard STDOUT & STDERR
echo
echo "Discarding STDOUT & STDERR:"
head -n3 /etc/passwd /fakefile &> /dev/null

# clean up
rm ${FILE} ${ERR_FILE} &> /dev/null
