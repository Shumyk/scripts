#!/bin/bash
#
# script generates a list of random passwords

# random number as a password
PASSWORD="${RANDOM}"
echo "${PASSWORD}"

# three random numbers together
PASSWORD="${RANDOM}${RANDOM}${RANDOM}"
echo "${PASSWORD}"

# use the current date/time as the basis for the password
PASSWORD=$(date +%s)
echo "${PASSWORD}"

# use nanoseconds to act as randomization
PASSWORD=$(date +%s%N)
echo "${PASSWORD}"


# a better password
PASSWORD=$(date +%s%N | sha256sum | head -c32)
echo "${PASSWORD}"

# an event better password
PASSWORD=$(date +%s%N${RANDOM}${RANDOM} | sha256sum | head -c48)
echo "${PASSWORD}"

# append a special character to the password
SPECIAL_CHARACTER=$(echo '!@#$%^&*()_+=' | fold -w1 | shuf | head -c1)
echo "${PASSWORD}${SPECIAL_CHARACTER}"
