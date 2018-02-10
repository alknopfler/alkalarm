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

echo "Starting the installation...Please press a key:"
read

echo "### Installing packages ... ###"
apt-get -y update
apt-get -y install golang python-pip python-dev pigpio python-pigpio python3-pigpio

echo "### Creating Folders ... ###"

create_folder $PROJECT_PATH
create_folder $TMP_PATH
create_folder $LOG_PATH

echo "### Copy the python script to the project directory ... ### "
cp _433.py $PROJECT_PATH/

echo "### Creating the password file with email smtp server password ... ### "

read -p "Introduce la password de la cuenta smtp: " password
echo $password>$PROJECT_PATH/.passSMTP

echo "### Enabling the PiGPIO daemon ###"
systemctl enable pigpiod


echo "### End Of Installation...SUCESS ###"
