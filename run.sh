#!/bin/bash
echo "Begin execution"
# [0-50]
echo $[$RANDOM%51]'%......'
# 如果是数据库模式, 还需要生成秘钥
if [ $(cat config.yaml | grep 'data_source:' | awk '{print $2}') == '1' ] && [ ! -d "secret/rsa" ]; then
  mkdir secret/rsa
  openssl genrsa -out secret/rsa/rsa_private_key.pem 1024
  openssl rsa -in secret/rsa/rsa_private_key.pem -pubout -out secret/rsa/rsa_public_key.pem
fi
sleep 1
# [50-89]
echo $[$RANDOM%40+50]'%......'
sleep 1
# 执行脚本
pid=$(lsof -i:4398 | awk '{print $2}' | grep '[0-9]')
# 如果已启动了就kill掉
if [ "$pid" == '' ]; then
  echo 'server not run'
else
  kill -9 $pid
  echo 'server running, kill '$pid
fi
go build main.go && go run main.go -log_dir=log -alsologtostderr > out.out 2>&1 &
# 防止rm 执行太快
sleep 1
# [90-99]
echo $[$RANDOM%10+90]'%......'
sleep 1
rm -f main
echo "100% run success!"