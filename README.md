# docker-test


logs example
```
ping-service-1  | Server starting on port 8080...
ping-service-1  | 2025/02/11 13:56:04 Successfully pinged http://docker-test-ping-service-1:8080/pings, status: 200
ping-service-1  | 2025/02/11 13:56:04 Successfully pinged http://docker-test-ping-service-2:8080/pings, status: 200
ping-service-2  | 2025/02/11 13:56:09 Successfully pinged http://docker-test-ping-service-2:8080/pings, status: 200
ping-service-2  | 2025/02/11 13:56:09 Successfully pinged http://docker-test-ping-service-1:8080/pings, status: 200
ping-service-1  | 2025/02/11 13:56:09 Successfully pinged http://docker-test-ping-service-2:8080/pings, status: 200
ping-service-1  | 2025/02/11 13:56:09 Successfully pinged http://docker-test-ping-service-1:8080/pings, status: 200
ping-service-2  | 2025/02/11 13:56:14 Successfully pinged http://docker-test-ping-service-2:8080/pings, status: 200
ping-service-2  | 2025/02/11 13:56:14 Successfully pinged http://docker-test-ping-service-1:8080/pings, status: 200
ping-service-1  | 2025/02/11 13:56:14 Successfully pinged http://docker-test-ping-service-2:8080/pings, status: 200
ping-service-1  | 2025/02/11 13:56:14 Successfully pinged http://docker-test-ping-service-1:8080/pings, status: 200
ping-service-2  | 2025/02/11 13:56:19 Successfully pinged http://docker-test-ping-service-2:8080/pings, status: 200
ping-service-2  | 2025/02/11 13:56:19 Successfully pinged http://docker-test-ping-service-1:8080/pings, status: 200
```