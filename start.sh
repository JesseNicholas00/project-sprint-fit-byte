pkill fit-byte.out
rm logs.txt
chmod +x fit-byte.out
nohup ./fit-byte.out > logs.txt &
echo application started