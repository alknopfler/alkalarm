#!/bin/bash

echo "   _               _
   / \  | | | __   / \  | | __ _ _ __ _ __ ___
  / _ \ | | |/ /  / _ \ | |/ _' | '__| '_ ' _ \
 / ___ \| |   <  / ___ \| | (_| | |  | | | | | |
/_/   \_\_|_|\_\/_/   \_\_|\__,_|_|  |_| |_| |_|
                                                "

PROJECT_PATH="/opt/alkalarm"
LOG_PATH="/var/log/alkalarm"
TMP_PATH="/var/tmp/alkalarm"


function create_folder(){
	if [ -d $1 ];
	then
		echo "$1 exists!"
		return 0
	fi
	mkdir $1
	if [ $? -ne 0 ];
	then
		echo "Couldn't create $1"
		exit 3
	else
		echo "Created $1"
	fi

	chown root:root $1
	if [ $? -ne 0 ];
	then
		echo "Couldn't change user and group of $1 to the specified user and group (root, root)!"
		exit 4
	else
		echo "Changed user and group of $1 to root, root"
	fi
}


##### main #####

tput setaf 2;echo "Starting the installation...Please press a key:"
read

tput setaf 2;echo "### Installing packages ... ###"
apt-get -y update
apt-get -y install golang python-pip python-dev pigpio python-pigpio python3-pigpio

tput setaf 2;echo "### Creating Folders ... ###"
create_folder $PROJECT_PATH
create_folder $TMP_PATH
create_folder $LOG_PATH

tput setaf 2;echo "### Copy the python script and binary to the project directory ... ### "
go install alkalarm.go
cp _433.py /root/bin/alkalarm $PROJECT_PATH/

tput setaf 2;echo "### Installing web server and copy the binary to the project directory ... ###"
go install webinterface/webserver.go
cp /root/bin/webserver $PROJECT_PATH/


tput setaf 2;echo "### Creating the password file with email smtp server password ... ### "
read -p "Introduce la password de la cuenta smtp: " password
echo $password>$PROJECT_PATH/.passSMTP

tput setaf 2;echo "### Enabling the PiGPIO daemon ###"
systemctl enable pigpiod

tput setaf 2;echo "### Creating the file in rsyslog ###"
cp 30-alkalarm.conf /etc/rsyslog.d/
systemctl restart rsyslog

tput setaf 2;echo "### Creating AlkAlarm systemd service ###"
cp alkalarm.service /lib/systemd/system/.
chmod 755 /lib/systemd/system/alkalarm.service
systemctl enable alkalarm
systemctl start alkalarm

tput setaf 2;echo "### Creating webserver systemd service ###"
cp alkalarm.service /lib/systemd/system/.
chmod 755 /lib/systemd/system/alkalarm.service
systemctl enable alkalarm
systemctl start alkalarm

tput setaf 2;echo "### End Of Installation...SUCCESS ###"
