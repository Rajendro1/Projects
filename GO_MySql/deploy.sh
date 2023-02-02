#! /bin/bash

cd ..
# scp -r Thirdessential_GO root@194.163.40.86://usr/local/lsws/dev.wikibedtimestories.com/html/webservices/
# scp -r Thirdessential_GO/includes/Db_Operation.go root@194.163.40.86://usr/local/lsws/dev.wikibedtimestories.com/html/webservices/Thirdessential_GO/includes/Db_Operation.go
scp -r Thirdessential_GO/api/product.go root@194.163.40.86://usr/local/lsws/dev.wikibedtimestories.com/html/webservices/Thirdessential_GO/api/product.go
# scp -r Thirdessential_GO/includes/Config.go root@194.163.40.86://usr/local/lsws/dev.wikibedtimestories.com/html/webservices/Thirdessential_GO/includes/Config.go
# ssh root@194.163.40.86 sudo 'systemctl restart thirdessential.service'

