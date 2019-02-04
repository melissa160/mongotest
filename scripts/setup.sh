mkdir /root/.ssh/
cp -r ci_config/* /root/.ssh/
touch /root/.ssh/id_rsa
chmod 600 /root/.ssh/id_rsa
touch /root/.ssh/known_host
ssh-keyscan github.com >> /root/.ssh/known_hosts
rm -rf /root/.ssh/*
