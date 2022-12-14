## PortMonitor

It requires HackerTarget API and monitor port using cronjob, and it requires mongodb to store data and slack webhook.

### MongoDB Compose file

This docker-compose file taken from [dockerhub](https://hub.docker.com/_/mongo)

```yaml
# Use root/example as user/password credentials
version: '3.1'

services:

  mongo:
    image: mongo
    restart: always
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: example
    ports:
      - 27017:27017

  mongo-express:
    image: mongo-express
    restart: always
    ports:
      - 8081:8081
    environment:
      ME_CONFIG_MONGODB_ADMINUSERNAME: root
      ME_CONFIG_MONGODB_ADMINPASSWORD: example
      ME_CONFIG_MONGODB_URL: mongodb://root:example@mongo:27017/
```

### Environment

Change HackerTarget API also mongodb URI and finally slack webhook. Create folder on `mkdir $HOME/.portmonitor/` and
then create a file `touch $HOME/.portmonitor/.env` follow the example env and change with yours.

```dotenv
HACKERTARGET_API=7xxxxxsomerandomstringsd1c
MONGODB_URI=mongodb://root:example@localhost:27017/
SLACK_WEBHOOK=https://hooks.slack.com/services/xxxx/xxx/xxxxx
```
