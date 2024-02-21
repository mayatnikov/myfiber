#!/bin/bash

for k in {1..1000}; do
 for host in 172.20.0.98;  do 
   echo HOST $host 
   for i in {1..20}; do
     echo -n $i ")"
     curl -XGET $host':8008/nm/Vitaly'
     curl -XGET $host':8008/json1'
     curl -XGET $host':8008/json2'
     curl -XGET $host':8008/json3'
     echo
   done
   echo '------------ // -------------'
 done
done
