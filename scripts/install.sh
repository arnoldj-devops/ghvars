#!/bin/bash
if which ghvars &>/dev/null
then
{    
    which ghvars | xargs sudo rm -rf 
}
fi
LATEST_RELEASE=$(curl -L -s -H 'Accept: application/json' https://github.com/arnoldj-devops/ghvars/releases/latest)
LATEST_VERSION_TAG=$(echo $LATEST_RELEASE | sed -e 's/.*"tag_name":"\([^"]*\)".*/\1/')
LATEST_VERSION="${LATEST_VERSION_TAG:1}"
ARTIFACT_URL=https://github.com/arnoldj-devops/ghvars/releases/download/${LATEST_VERSION_TAG}/ghvars_${LATEST_VERSION}_linux_amd64.tar.gz
wget -qc  $ARTIFACT_URL -P /tmp && sudo tar -xvf /tmp/ghvars_${LATEST_VERSION}_linux_amd64.tar.gz -C /usr/local/bin/ ghvars >/dev/null 2>&1
ghvars --help
