# mcping

An easy way to know if a memcached server is up

memping (http://libmemcached.org) keeps running forever if the remote server cannot be reached, I wanted a simple tool to use in scripts and know immediatly if a memcached server is up or not, introducing mcping.

memcached server up and reachable

    mcping 127.0.0.1:11211
    echo $?
    0

memcached server down or not reachable

    pkill memcached
    mcping 127.0.0.1:11211
    2016/08/07 12:58:35 dial tcp 127.0.0.1:11211: getsockopt: connection refused
    echo $?
    1
