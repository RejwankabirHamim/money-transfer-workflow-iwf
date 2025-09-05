# How to run 

## 1. start a iWF server following the instruction:
 ```bash
   docker pull iworkflowio/iwf-server-lite:v1.18.3 && docker run -p 8801:8801 -p 7233:7233 -p 8233:8233 -e AUTO_FIX_WORKER_URL=host.docker.internal --add-host host.docker.internal:host-gateway -it iworkflowio/iwf-server-lite:v1.18.3
```
This by default will run Temporal server with it, again:

* IWF service: http://localhost:8801/
* Temporal WebUI: http://localhost:8233/
* Temporal service: localhost:7233

## 2. build and run this project 
 ```bash 
   make bins && ./iwf-samples start
 ```
## 3. start a workflow:
http://localhost:8803/moneytransfer/start?fromAccount=test1&toAccount=test2&amount=100&notes=hello

## 4. watch in WebUI
http://localhost:8233/namespaces/default/workflows

