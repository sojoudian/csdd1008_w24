 docker build -t maziar/torontotime:v01 .
 docker push  maziar/torontotime:v01
 docker run -p 8014:8014 maziar/torontotime:v01
