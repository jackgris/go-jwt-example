# Go JWT Example
This example will show you how you can create an JWT service using Fiber. (HS256)

## After running the server you can tested with CURL:

_Login using username and password to retrieve a token._
```
curl --data "user=john&pass=doe" http://localhost:3000/login
```
_Response_
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjg0NzEwNTMxLCJuYW1lIjoiSm9obiBEb2UifQ.kbZn3zXPIdZt4hJIb3EvodVh3UO3jdTyPgxK8huxDFM"
}
```


_Request a restricted resource using the token in Authorization request header._
```
curl localhost:3000/restricted -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjg0NzEwNTMxLCJuYW1lIjoiSm9obiBEb2UifQ.kbZn3zXPIdZt4hJIb3EvodVh3UO3jdTyPgxK8huxDFM"
```
_Response_
```
Welcome John Doe
```
## Using Makefile

To make everything easy, you can use the make commands to run the server and the client to see how this work. In the folder of the project:

First run the server with:
```
make run-server
```

And then you can run the times you want the client:

```
make run-client
```

After all, you can use this command for cleaning purposes:

```
make clean
```

Note: The command in this make file run in a Linux terminal.

This example was written for this post on my blog: ['Building an Authorization Service With Fiber Using JWT'](https://jackgris.github.io/goscrapy-blog/post/building-an-authorization-service-with-fiber-using-jwt/)
