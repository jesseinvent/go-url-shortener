## Clone repo

`$ git clone https://github.com/jesseinvent/go-url-shortener`

## Run containers

`$ docker-compose up -d --build `

## Endpoints

### Create Shortened link

`POST /api/v1/create_link `

```
{
    "url": "https://youtu.be/03VxdOzPFSQ",
    "alias": "youtube_vid"
    "expiry": 24
}
```
