<p align="center">
    <a href="https://tech.wildberries.ru" target="_blank">
        <img src="wb.png" alt="wb">
    </a>
</p>

## WB Tech: level # 0 (Golang)

### Программа курса

- Go
- PostgreSQL
- NATS Streaming
- Gin Web Framework

### Получение данных о заказе

- Request URL: `http://localhost:8080/api/order-models/b563feb7b2b84b6test`
- Method: `GET`
- Path: `/order-models/{id}`
- Headers: `Accept:application/json, Content-Type:application/json`
- Status Code: `200`

### Тесты WRK и Vegeta
```
go-wrk -c 100 -d 5 http://localhost:8080/api/order-models/b563feb7b2b84b6test
Running 5s test @ http://localhost:8080/api/order-models/b563feb7b2b84b6test
  100 goroutine(s) running concurrently
105115 requests in 4.995706843s, 95.33MB read
Requests/sec:           21041.07
Transfer/sec:           19.08MB
Avg Req Time:           4.75261ms
Fastest Request:        54.513µs
Slowest Request:        563.536158ms
Number of Errors:       0
```

```
echo "GET http://localhost:8080/api/order-models/b563feb7b2b84b6test" | vegeta attack -duration=20s -rate=100 | tee results.bin | vegeta report
Requests      [total, rate, throughput]         2000, 100.05, 100.05
Duration      [total, attack, wait]             19.991s, 19.991s, 203.014µs
Latencies     [min, mean, 50, 90, 95, 99, max]  142.81µs, 323.636µs, 315.142µs, 394.235µs, 433.521µs, 657.348µs, 1.233ms
Bytes In      [total, mean]                     1688000, 844.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:2000  
Error Set:
```
