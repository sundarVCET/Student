step :1 Find the process ID:

lsof -i :PORT_NUMBER
lsof -i :8080

step:2 kill PID (process id)
kill -9 PID


