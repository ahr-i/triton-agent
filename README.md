# Triton Agent
This is the agent code managing the Triton container.  

## 1. Docker Start
### 1.1 Download
```
git clone https://github.com/ahr-i/triton-agent.git
```

### 1.2 Setting
```
cd triton-agent
vim setting/setting.go
```
Modify the contents of the file.   
```
package setting

/* ----- Server Setting ----- */
const ServerPort string = "7000" // Edit this

/* ----- Triton Server Setting ----- */
const TritonUrl string = "localhost:8000" // Edit this

// If you are not using a scheduler, change the 'SchedulerActive' variable to false.
const SchedulerActive bool = false           // Edit this
const SchedulerUrl string = "localhost:8000" // Edit this
```
It is advisable to set Triton with a fixed IP.   
The triton-agent and the Triton container must be on the same network.   

To create a Docker Network, use the following command.
```
docker network create --subnet=100.0.0.0/24 triton
```
To run the Triton container with a fixed IP, add the following option when launching the Triton container.
```
--network triton --ip 100.0.0.2
```

### 1.3 Build
```
docker build -t triton-agent .
```

### 1.4 Run
```
docker run -it --rm --name triton-agent --network triton -p 7000:7000 triton-agent
```
