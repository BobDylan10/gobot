# create random password
#PASSWDDB="$(openssl rand -base64 12)"
PASSWDDB="gobot"

MAINDB="gobot_db"
USER_NAME="gobot"

mysql -e "CREATE DATABASE ${MAINDB} DEFAULT CHARACTER SET utf8;"
mysql -e "CREATE USER '${USER_NAME}'@'localhost' IDENTIFIED BY '${PASSWDDB}';"
mysql -e "GRANT ALL PRIVILEGES ON ${MAINDB}.* TO '${USER_NAME}'@'localhost';"
mysql -e "FLUSH PRIVILEGES;"