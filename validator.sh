#!/usr/bin/env bash

#cd ../../../fca/cbt
env GOOS="linux" GOARCH="amd64" go build -v
scp spring_block javaxu@ecru.cs.ucl.ac.uk
sshpass -p "92072910" scp spring_block javaxu@ecru.cs.ucl.ac.uk:/cs/research/external/home/javaxu/.sbl
sshpass -p "92072910" ssh javaxu@ecru.cs.ucl.ac.uk /bin/bash << EOF
cd .sbl
ls -al
./spring_block

EOF