## Blacklist word matcher

Golang serverless function to check article submissions against a word blacklist and approve/deny accordingly. This was built for the ability to have a good performance when using blacklists consisting of thousands of words.

Cloudflare's Go implemention of Aho-Corasick was also tested but showed no benefits against the standard library with blacklist consisting of 10.000+ keywords. The code is commented out, if someone would like to use it.

### Usage

Clone repo

`git clone git@github.com:MalpraveCorp/blacklist-matcher.git`

Build project

`make`

Deploy

`sls deploy`

Load test

`k6 run loadtest.js`

API

Send a POST request with the article in the body

`curl https://$ENDPOINT/dev/blacklist-matcher -d "article contents"
`
### Performance
Using the blacklist provided in this repo, here's the load testing results showing a p95 performance under 100ms (skipping cold starts). 

```

running (0m40.0s), 00/30 VUs, 8180 complete and 0 interrupted iterations
default ✓ [======================================] 00/30 VUs  40s


    ✓ status was 400

    checks.....................: 100.00% ✓ 8180 ✗ 0   
    data_received..............: 4.0 MB  100 kB/s
    data_sent..................: 28 MB   698 kB/s
    http_req_blocked...........: avg=180.09µs min=273ns   med=557ns   max=110.96ms p(90)=823ns    p(95)=930ns   
    http_req_connecting........: avg=78.66µs  min=0s      med=0s      max=26.78ms  p(90)=0s       p(95)=0s      
    http_req_duration..........: avg=71.5ms   min=58.07ms med=70.83ms max=224.37ms p(90)=78.26ms  p(95)=82.87ms 
    http_req_receiving.........: avg=129.26µs min=31.04µs med=73.24µs max=11.91ms  p(90)=176.28µs p(95)=452.24µs
    http_req_sending...........: avg=96.07µs  min=42.32µs med=82.3µs  max=21.64ms  p(90)=117.86µs p(95)=143.34µs
    http_req_tls_handshaking...: avg=95.9µs   min=0s      med=0s      max=54.72ms  p(90)=0s       p(95)=0s      
    http_req_waiting...........: avg=71.27ms  min=57.31ms med=70.63ms max=224.16ms p(90)=77.96ms  p(95)=82.58ms 
    http_reqs..................: 8180    204.317965/s
    iteration_duration.........: avg=72.84ms  min=58.98ms med=71.89ms max=225.42ms p(90)=79.7ms   p(95)=84.88ms 
    iterations.................: 8180    204.317965/s
    vus........................: 1       min=1  max=29
    vus_max....................: 30      min=30 max=30
```
### License
MIT
