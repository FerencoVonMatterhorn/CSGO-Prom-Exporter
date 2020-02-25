<img align="center" src="https://lh3.googleusercontent.com/proxy/CRbywP1kEljBZjCc7E03iEPXeONbLRopXYdCqBBg1k26sCgGspOpuhCNYSxPDf05GYTRKETMQndRBpbNZFTpfGHpsoKp75mt9Bh5bjvtZK-tPwyZiSOs34I_JAo" />

# CSGO-Prometheus-Exporter

The CSGO-Exporter is a Prometheus Exporter, written in golang, which lets you see your CSGO Stats, that you wouldnt normally see.

![img not available](https://raw.githubusercontent.com/ferencovonmatterhorn/csgo-prom-exporter/master/grafana-dashboard.png)

## Installation
To run the whole stack, just change your the credentials in the docker-compose file `services.csgo-exporter.environment:` and then run the file by typing: 
```Dockerfile
docker-compose up -d
```

 and it will automatically fetch and install Prometheus, Grafana, Node-Exporter, and the CSGO-Exporter.

```Dockerfile
version: '3.3'

volumes:
  prometheus:
  grafana:

services:
  grafana:
    image: grafana/grafana:latest
    ports:
       - 3000:3000
    volumes:
      - grafana:/var/lib/grafana
    restart: always
 
  prometheus:
    image: prom/prometheus:latest
    ports:
      - 9090:9090
    volumes:
      - ./prometheus/:/etc/prometheus
      - prometheus:/prometheus
    command:
      - '--config.file=/etc/prometheus/prometheus.yml'
      - '--storage.tsdb.path=/prometheus'
    restart: always

  node-exporter:
    image: prom/node-exporter:latest
    ports:
      - 9200:9200
    volumes:
      - /proc:/host/proc:ro
      - /sys:/host/sys:ro
      - /:/rootfs:ro
    command:
      - '--path.procfs=/host/proc'
      - '--path.rootfs=/rootfs'
      - '--path.sysfs=/host/sys'
      - '--collector.filesystem.ignored-mount-points=^/(sys|proc|dev|host|etc)($$|/)'
    restart: unless-stopped
 
  csgo-exporter:
    build: ./
    environment: 
      - EXPORTERPORT=9092           #Default Port -> prometheus/prometheus.yml
      - STEAMID=<YourSteamID>
      - STEAMKEY=<YourSteamKey>
    restart: unless-stopped
```

## How it Works

The Exporter makes simple API calls to the Steam-API Servers, takes the information and parses it into readable Metrics for Prometheus.


## TODO

- [x] Use smaller Alpine Image
- [ ] Import more Metrics
- [x] Use go binary for smaller Image Size
- [ ] Add Env for Loglevel


## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

