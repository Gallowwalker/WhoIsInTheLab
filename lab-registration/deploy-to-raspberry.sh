#!/bin/bash
# Simple deploy script for registration server

# If no arguments specified print usage and exit
[ $# -eq 0 ] && { echo "Usage: $0 [remote_ip]"; exit 1; }

echo "Updating js deps from bower"
cd ../ && bower install

echo "Cross compiling GO code for arm" 
cd lab-registration
go clean
GOARCH=arm GOOS=linux go build
echo "Finished compiling GO code"

echo "Archiving server binary and client side code"
tar -cf lab-registration.tar ./lab-registration config public

echo "Enter password to upload archive to remote host"
scp ./lab-registration.tar hackafe@$1:/home/hackafe/ 


echo "Enter password again to restart the server"
ssh hackafe@$1 "bash -s" << EOF 
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
