wrk.method = "POST"
wrk.headers["Content-Type"] = "application/application/json"
wrk.body = '{"version": "v1.0.0","token": "token","start_stop": true,"data": [1, 2, 3, 4]}'

## wrk压测
## wrk -t4 -c2000 -d60s -T5s --script=./wrk_post.lua http://127.0.0.1:8090/