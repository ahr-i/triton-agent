# Triton Agent
This is the agent code managing the Triton container.  

## 1. Docker Start
### 1.1 Clone
```
git clone https://github.com/ahr-i/triton-agent.git
```

### 1.2 build
```
cd triton-agent
docker build -t triton-agent .
```

### 1.3 setting
```
vim setting/setting.go
```
Modify the contents of the file.   
```
package setting

/* ----- Server Setting ----- */
const ServerPort string = "1000" // Edit this

/* ----- Triton Server Setting ----- */
const TritonUrl string = "localhost:2000"    // Edit this
const SchedulerUrl string = "localhost:8000" // Edit this
```
It is advisable to set Triton with a fixed IP.   
The triton-agent and the Triton container must be on the same network.   

To create a Docker Network, use the following command.
```
docker network create --subnet=100.0.0.0/24 test
```
To run the Triton container with a fixed IP, add the following option when launching the Triton container.
```
--network test --ip 100.0.0.2
```

### 1.4 Run
```
docker run -it --rm --network test -p 1000 triton-agent
```
