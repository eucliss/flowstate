Flowstate is a monitoring tool for monitoring the status of automation flows.


How to run:
Start go server:
```
go run flowstate.go
```
Start frontend:
```
npm run dev
```

Start openobserve:
```
ZO_ROOT_USER_EMAIL="email" ZO_ROOT_USER_PASSWORD="pass" ./openobserve
```

Start redis:
```
redis-stack-server
```

Steps:
1. Setup a Go backend server
    - Webhook reciever to recieve 


ToDo:
- handle inbound/outbound edges being deleted when new nodes are deleted.
- Change Node ID structure
- Setup log stuff
