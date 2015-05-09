#!/bin/bash
set -e

# Test for presence of python 2.7.3
# Test for presence of curl

# Test that output folder exists

# TODO: Replace with env vars
USER_AGENT="kdeloach@azavea.com"
OAUTH_TOKEN="813a87727d1bc0862cd3b1908fb1060fe4b2f3cf"

# Generate list of repos in Waffle with:
# $('[bo-bind="source.repoPath"]').map(function(){ return $(this).text(); }).toArray().join("\n")

#REPOS=(
#azavea/backbone.hashmodels
#azavea/civic-apps
#azavea/django-queryset-csv
#azavea/django-tinsel
#azavea/dor-parcel-explorer
#azavea/dor-philly-history-blog
#azavea/lr-common
#azavea/model-my-watershed
#azavea/nyc-trees
#azavea/oit-ulrs
#azavea/OpenTreeMap-iOS-skin
#azavea/pwd-stormwater-allocation
#azavea/pwd-stormwater-interactive
#azavea/pwd-waterworks-revealed
#azavea/stanford-campus-map
#azavea/tr-55
#azavea/usace-wisdm
#azavea/usace-wisdm-filter-viewer
#azavea/usace-wisdm-filter-viewer-data
#azavea/usace-wisdm-symbolizer
#azavea/usace-wisdm-symbolizer-data
#CoastalResilienceNetwork/GeositeFramework
#OpenTreeMap/clients
#OpenTreeMap/cloudbuild
#OpenTreeMap/ecobenefits
#OpenTreeMap/OpenTreeMap
#OpenTreeMap/OpenTreeMap-Android
#OpenTreeMap/OpenTreeMap-iOS
#OpenTreeMap/OpenTreeMap-Modeling
#OpenTreeMap/opentreemap.github.com
#OpenTreeMap/otm-mobile-skins
#OpenTreeMap/otm-wordpress
#OpenTreeMap/otm-wordpress-release-scripts
#OpenTreeMap/OTM2
#OpenTreeMap/otm2-addons
#OpenTreeMap/OTM2-tiler
#OpenTreeMap/otm2-vagrant
#)

REPOS=(
azavea/civic-apps
azavea/model-my-watershed
azavea/nyc-trees
)

QUERY_ARGS="milestone=*&state=open"

main() {
    for repo in ${REPOS[*]}; do
        url="https://api.github.com/repos/${repo}/issues?${QUERY_ARGS}"
        file="output/${repo/\//-}-issues.json"
        curl -v --silent \
             -H "Authorization: token ${OAUTH_TOKEN}" \
             -H "User-Agent: ${USER_AGENT}" \
             ${url} 2>/dev/null > ${file}
        echo "Updated ${file}"
        sleep 0.5
    done
    exit 0
}

main
