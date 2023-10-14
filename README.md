![shadowtester](https://github.com/aminGhafoory/shadowsocksTester/assets/74732606/053bf0f7-7a05-4402-ba03-946b605826c5)
# shadowsocksTester

this little program fetches shadowsockses from diffrent subscriptions and test them and rank them so you can use them more conveniently
you need to edit .env file based of your own database and your preferd port and you should run the migrations on your database
### add subscription link to subscription pool
```
curl  -X POST \
  'http://localhost:3000/sub' \
  --header 'Content-Type: application/json' \
  --data-raw '{"URL":"https://raw.githubusercontent.com/mahdibland/ShadowsocksAggregator/master/sub/splitted/ss.txt"}'
```
### get all subcriptions in sub pool
```
curl  -X GET 'http://localhost:3000/sub' 
```


### get all fethced shadowsockses
```
curl  -X GET 'http://localhost:3000/ss' 
```

### get best shadowsockses
```
curl  -X GET 'http://localhost:3000/best'
```

## TODO
  #### -- embeded database migrations
  
  #### -- Dockerize the whole app
