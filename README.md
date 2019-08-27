# httpsproxy

This is a simple proxy which listens on port 443 and forwards the request to
a given url (default is http://localhost) The certificates are selfsigned and
built into the command. 

I build this to allow serving https then using php:s built in webserver.
You can start serving with php like this

```
php -S localhost:8080
```

and then in another terminal run

```
httpsproxy http://localhost:8080
```

all requests to https://localhost will be forwarded to http://localhost:8080. If you want to simulate an address could you put something like this in your /etc/hosts
```
127.0.0.1 www.test.com
```


## Usage

httpsproxy [url]

## 