#!/bin/bash

# TEST COLORS
GR='\033[1;92m' NC='\033[0m'
# CONSTANTS
GCR_URL="us.gcr.io/<your-url>"
DOMAIN="your-domain"
STATEFULSET_MS="your-microservice"
PREVIOUS_IMAGE="/tmp/kdeploy"
# PARAMETERS
# microservice            show images
MS=${1}                   OCCURRENCES=${2:-20}


watchPods() {
    watch -n1 "kubectl get pods | grep ${MS}"
}

kubectlSetImage() {
    kubectl set image "${TYPE}/${NAMESPACE}-${MS}" "${MS}=${IMAGE}"
    printf "updated %s%s%s %s with image %s%s%s\n" "${GR}" "${MS}" "${NC}" "${TYPE}" "${GR}" "${IMAGE}" "${NC}"
}

deployPrevious() {
    [[ ! -f ${PREVIOUS_IMAGE} ]] && echo "No previous deployment" >&2 && return 1
    [[ "$(grep -vc ^$ < "${PREVIOUS_IMAGE}")" != 4 ]] && echo "Previous deployment file corrupted: ${PREVIOUS_IMAGE}" >&2 && return 1

    TYPE=$(sed '1q;d' ${PREVIOUS_IMAGE})
    NAMESPACE=$(sed '2q;d' ${PREVIOUS_IMAGE})
    MS=$(sed '3q;d' ${PREVIOUS_IMAGE})
    IMAGE=$(sed '4q;d' ${PREVIOUS_IMAGE})
    rm ${PREVIOUS_IMAGE}

    kubectlSetImage
    watchPods
}

savePreviousDeployment() {
    echo "${TYPE}" > ${PREVIOUS_IMAGE}
    echo "${NAMESPACE}"; echo "${MS}"; echo "${CURRENT_IMAGE}" >> ${PREVIOUS_IMAGE}
}

if [ "${1}" = "previous" ]; then deployPrevious; return 0; fi

# SCRIPT BEGINNING
# fetch all metadata
NAMESPACE=$(kubectl config view --minify --output 'jsonpath={..namespace}'; echo)
TYPE=$(if [[ "${MS}" = "${STATEFULSET_MS}" ]]; then echo "statefulset"; else echo "deployment"; fi)
CURRENT_IMAGE=$(kubectl get "${TYPE}" "${NAMESPACE}-${MS}" -o jsonpath="{..image}" | tr -s "[:space:]" "\n" | grep "${DOMAIN}-${MS}")

# save current deployment metadata for 'kdeploy previous' command
savePreviousDeployment
printf "currently deployed image: %s%s%s\n\n" "${GR}" "${CURRENT_IMAGE}" "${NC}"

IMAGES_LIST=$(gcloud alpha container images list-tags "${GCR_URL}${MS}" \
                --show-occurrences-from="${OCCURRENCES}" --format="value[separator=|](timestamp, digest, tags)")

SELECTED_IMAGE_TMP_FILE="tmp.file.selected.image"
# runs go to interactively select image to deploy
select_image_prompt ${SELECTED_IMAGE_TMP_FILE} $(echo ${IMAGES_LIST})
SELECTED_IMAGE_INFO=($(cat ${SELECTED_IMAGE_TMP_FILE}; echo))
rm ${SELECTED_IMAGE_TMP_FILE}

# finish script if nothing selected
if [ ${#SELECTED_IMAGE_INFO[@]} -eq 0 ]; then return; fi

SHORT_DIGEST=${SELECTED_IMAGE_INFO[2]}
TAG=${SELECTED_IMAGE_INFO[3]}

FULL_DIGEST=$(gcloud container images describe "${GCR_URL}${MS}@sha256:${SHORT_DIGEST}" \
                    --format="value(image_summary.digest)")

# if tag absent semicolon should be omitted
OPTIONAL_TAG_WITH_SEMICOLON=$(if [[ ${TAG} ]]; then echo ":${TAG}"; fi)
IMAGE="${GCR_URL}${MS}${OPTIONAL_TAG_WITH_SEMICOLON}@${FULL_DIGEST}"

kubectlSetImage
watchPods
