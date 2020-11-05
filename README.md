# GoRailwire


Bring the service up
```
docker-compose up
```

Try the save plan functionality
```
cat SavePlan.json | grpcurl -d @ -plaintext localhost:9092 Plans.SavePlan
```


Try the get plan functionality
```
cat GetKeralaPlan.json | grpcurl -d @ -plaintext localhost:9092 Plans.GetPlan
```
