#!/bin/bash

LOG_FILE="/home/photo/logs/cronenable"

# Remove a certain cron job (requires sudo)
#sudo crontab -l | grep -v -F '/media/usbHDD1/deps/archway-main/contrib/localnet/opt/liquidstake.sh' | sudo crontab -u photo -r
#sudo crontab -l | grep -v -F '/bin/date' | sudo crontab -u photo -r

# Write out current crontab
#echo "Checking for existing crontab for photo" >> $LOG_FILE
#if sudo crontab -u photo -l > /dev/null 2>&1; then
#    echo "Writing out current crontab" >> $LOG_FILE
#    sudo crontab -u photo -l > mycron
#else
#    echo "No crontab for photo" >> $LOG_FILE
#   touch mycron
#fi

#echo "Changing owner of log file to current user" >> $LOG_FILE
#LOG_FILE1="/var/log/cron_$(date +\%Y\%m\%d\%H\%M\%S).log"
touch $LOG_FILE
#sudo chown $(whoami):$(whoami) $LOG_FILE1

# Echo new cron into cron file
echo "enable cron" >> $LOG_FILE
#echo "enable cron" >> mycron
#echo "*/2 * * * * /bin/date >> /home/photo/myscript1.log" >> mycron

# Install new cron file
#echo "Installing new cron file" >> $LOG_FILE
#sudo crontab -u photo mycron

# Remove temporary file
#echo "Removing temporary cron file" >> $LOG_FILE


# Make script executable
#echo "Making script executable" >> $LOG_FILE
#sudo chmod +x /media/usbHDD1/deps/archway-main/contrib/localnet/opt/liquidstake.sh

# Change owner of log file to current user

