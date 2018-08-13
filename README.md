```bash
git clone https://github.com/andrewdruzhinin/travel-cost-app.git
```

Setup GOOGLEMAPS_API_KEY=YOUR_KEY at .env

```shell
docker-compose build
docker-compose up
```

```shell
curl --request GET \
  --url 'http://localhost:8080/trip/price?=' \
  --header 'content-type: application/json' \
  --data '{
	"start_point" : {
		"lat" : 53.277148,
		"long": 50.237069
	},
	"end_point" : {
		"lat" : 53.217118,
		"long": 50.188692
	}
}'
```