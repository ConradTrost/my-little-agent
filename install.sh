#!/bin/bash

# create user
sudo useradd -r -U my-little-agent

# make sure folders don't exist already
sudo rm -rf /usr/bin/my-little-agent
sudo rm -rf /etc/systemd/system/my-little-agent.service


# extract the tarball and move the files to the appropriate locations
tar xvzf /tmp/my-little-agent.tar.gz
rm /tmp/my-little-agent.tar.gz
sudo mkdir /usr/bin/my-little-agent
sudo mv /tmp/my-little-agent/main /usr/bin/my-little-agent/main
sudo mv /tmp/my-little-agent/my-little-agent.service /etc/systemd/system/my-little-agent.service
rm -rf /tmp/my-little-agent

# start the service
sudo systemctl daemon-reload
sudo systemctl enable my-little-agent.service
sudo systemctl is-enabled my-little-agent.service


# install script
# scp -i ~/.ssh/conrad-agent-kp.pem install.sh admin@172.30.49.44:/home/admin/my-little-agent-install.sh


# executable
# scp -i ~/.ssh/conrad-agent-kp.pem my_little_agent.tar.gz admin@172.30.49.44:/tmp/my-little-agent.tar.gz