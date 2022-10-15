#!/bin/bash
#
# this script demonstrates the case statement

# to much clutter with if-else
inconvenientIfElseConstruct() {
  if [[ "${1}" = 'start' ]]
  then
    echo 'Starting.'
  elif [[ "${1}" = 'stop' ]]
  then
    echo 'stopping'
  elif [[ "${1}" = 'status' ]]
  then
    echo 'status: '
  else
    echo 'supply a valid option.' >&2
    exit 1
  fi
}

verboseCase() {
  case "${1}" in
    start)
      echo 'starting'
      ;;
    stop)
      echo 'stopping'
      ;;
    status|state|--status|--state)
      echo 'status'
      ;;
    *)
      echo 'supply a valid option' >&2
      exit 1
      ;;
  esac
}

case "${1}" in
	start) echo 'starting' ;;
	stop) echo 'stopping' ;;
	status|state|--status|--state) echo 'status' ;;
	*)
		echo 'supply a valid option' >&2
		exit 1
		;;
esac
