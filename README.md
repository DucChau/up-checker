# Overview

    Simple service to return the status of a URL (listens on port 8000)

# Build 

`go binary`
    
```bash 
go build -o up-checker && ./up-checker
```

`docker` 

```bash
docker build -t up-checker
```

# Run 

`go binary`

```bash 
./up-checker
```

`docker`
    
```bash
docker run -p 8000:8000 -it up-checker
```

# Request (POST)

```json
{  
   "url": "https://www.omaze.com"
}
```

# Response

```json
{
    "time": 1588291315,
    "url": "https://www.omaze.com",
    "code": 200,
    "status": "200 OK"
}
```

# Test (build docker, start docker, post to server with curl)
```bash
docker build -t up-checker . && 
  docker run -d -p 8000:8000 up-checker && 
  curl -d '{"url": "https://www.omaze.com"}' -H "Content-Type: application/json" -X POST http://localhost:8000/v1/health && 
  docker ps
```
