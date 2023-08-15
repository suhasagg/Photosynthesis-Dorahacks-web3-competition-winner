#!/bin/bash

LOG_FILE="/home/photo/myscript.log"

# Write out current crontab
echo "Checking for existing crontab for photo" >> $LOG_FILE
if crontab -u photo -l > /dev/null 2>&1; then
    echo "Writing out current crontab" >> $LOG_FILE
    crontab -u photo -l > mycron
else
    echo "No crontab for photo" >> $LOG_FILE
    touch mycron
fi

echo "Changing owner of log file to current user" >> $LOG_FILE
LOG_FILE1="/var/log/cron_$(date +\%Y\%m\%d\%H\%M\%S).log"
touch $LOG_FILE1
chown $(whoami):$(whoami) $LOG_FILE1

# Echo new cron into cron file
echo "Adding new cron job" >> $LOG_FILE
echo "* * * * * /bin/sh /media/usbHDD1/deps/archway-main/contrib/localnet/opt/liquidstake.sh >> $LOG_FILE1 2>&1" >> mycron

# Install new cron file
echo "Installing new cron file" >> $LOG_FILE
crontab -u photo mycron

# Remove temporary file
echo "Removing temporary cron file" >> $LOG_FILE
rm mycron

# Make script executable
echo "Making script executable" >> $LOG_FILE
chmod +x /media/usbHDD1/deps/archway-main/contrib/localnet/opt/liquidstake.sh

# Change owner of log file to current user

