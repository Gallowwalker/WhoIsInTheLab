#!/bin/bash
# Simple deploy script for registration server

# If no arguments specified print usage and exit
[ $# -eq 0 ] && { echo "Usage: $0 [remote_ip] [ssh_port] [arch]"; exit 1; }

declare ssh_port=22
declare target_arch=386

if [ ${#2} != 0 ] 
then
	ssh_port=$2 
fi
if [ ${#3} != 0 ]
then 
	target_arch=$3
fi

echo "Updating js deps from bower"
cd ../ && bower install

echo "Cross compiling GO code for " $target_arch 
cd lab-registration
go clean
GOARCH=$target_arch GOOS=linux go build
echo "Finished compiling GO code"

echo "Archiving server binary and client side code"
tar -cf lab-registration.tar ./lab-registration config public

echo "Enter password to upload archive to remote host"
scp -P$ssh_port ./lab-registration.tar hackafe@$1:/home/hackafe/ 


echo "Enter password again to restart the server"
ssh hackafe@$1 -p $ssh_port "bash -s" << EOF 
			echo "Extracting server archive..."
			mkdir -p lab-registration 
			mv lab-registration.tar lab-registration 
			cd lab-registration
			tar -xf lab-registration.tar
			echo "Restarting server..."
			pkill -f "./lab-registration -config /aux0/WhoIsInTheLab/db.cfg"
			nohup ./lab-registration -config /aux0/WhoIsInTheLab/db.cfg > /dev/null 2>&1 &
EOF

go clean
