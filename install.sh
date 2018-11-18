#!/bin/bash

tput setaf 1;echo "
        _ _         _
       | | |       | |
   __ _| | | ____ _| | __ _ _ __ _ __ ___
  / _  | | |/ / _  | |/ _  | '__| '_   _ \
 | (_| | |   < (_| | | (_| | |  | | | | | |
  \__,_|_|_|\_\__,_|_|\__,_|_|  |_| |_| |_|
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
tput sgr0;read

tput setaf 2;echo
echo "### Installing packages ... ###"
tput sgr0;echo
apt-get -y update
apt-get -y install python-pip python-dev pigpio python-pigpio python3-pigpio
wget https://storage.googleapis.com/golang/go1.10.1.linux-armv6l.tar.gz
sudo tar -C /usr/local -xvf go1.10.1.linux-armv6l.tar.gz
ln -s /usr/local/go/bin/go /usr/bin/go


tput setaf 2;echo
echo "### Creating Folders ... ###"
tput sgr0;echo
create_folder $PROJECT_PATH
create_folder $TMP_PATH
create_folder $LOG_PATH

tput setaf 2;echo
echo "### Creating the email smtp server account ... ### "
tput sgr0;read -p "Introduce el usuario desde el que enviar notificaciones (pej. pepe@gmail.com) : " useraccount
echo -n $useraccount>$PROJECT_PATH/.userSMTP

tput setaf 2;echo
echo "### Creating the password file with email smtp server password ... ### "
tput sgr0;read -p "Introduce la password de la cuenta smtp: " password
echo $password>$PROJECT_PATH/.passSMTP

tput setaf 2;echo
echo "### Creating the accounts to use the alarm system ... ### "
tput sgr0;read -p "Introduce el listado de emails para usar la alarma (separado por comas): " listAccounts
echo -n $listAccounts>$PROJECT_PATH/.listACCESS

cp creds.json $PROJECT_PATH

tput setaf 2;echo
echo "### Creating the domain for the webserver ... ### "
tput sgr0;read -p "Introduce el host o la IP del host alkalarm: " HOST
cp -r webinterface $PROJECT_PATH
sed -i -- "s/{{HOST}}/$HOST/g" $PROJECT_PATH/webinterface/static/scripts/*.js

tput setaf 2;echo
echo "### Copy the python script and binary to the project directory ... ### "
tput sgr0;echo
go install alkalarm.go
cp _433.py /root/bin/alkalarm $PROJECT_PATH/

tput setaf 2;echo
echo "### Installing web server and copy the binary to the project directory ... ###"
tput sgr0;echo
go install webinterface/webserver.go

cp /root/bin/webserver $PROJECT_PATH/webinterface/

tput setaf 2;echo
echo "### Installing the discovery codes utility and copy the binary to the project directory ... ###"
tput sgr0;echo
go install scripts/discoverCodes.go
cp /root/bin/discoverCodes $PROJECT_PATH/

tput setaf 2;echo
echo "### Creating the file in rsyslog ###"
tput sgr0;echo
cp 30-alkalarm.conf /etc/rsyslog.d/
systemctl restart rsyslog

tput setaf 2;echo
echo "### Enabling the PiGPIO daemon ###"
tput sgr0;echo
systemctl enable pigpiod

tput setaf 2;echo
echo "### Creating and starting AlkAlarm systemd service ###"
tput sgr0;echo
cp alkalarm.service /lib/systemd/system/.
chmod 755 /lib/systemd/system/alkalarm.service
systemctl enable alkalarm
systemctl start alkalarm

tput setaf 2;echo
echo "### Creating and starting webserver systemd service ###"
tput sgr0;echo
cp alkalarm-webserver.service /lib/systemd/system/.
chmod 755 /lib/systemd/system/alkalarm-webserver.service
systemctl enable alkalarm-webserver
systemctl start alkalarm-webserver

tput setaf 2;echo
echo "### End Of Installation...SUCCESS ###"
tput sgr0;echo