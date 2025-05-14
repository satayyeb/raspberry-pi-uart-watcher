## Raspberry Pi UART watcher and error detector

![](https://images.ctfassets.net/3prze68gbwl1/asset-17suaysk1qa1hys/a23693cc21488367ff64bc1e7822d370/raspberry-pi-motion-sensor-detector.png)


### How to deploy?
1- customize network ENVs in the `Makefile`.

2- run:
```shell
  make deploy
```

3- You can see errors log by running:
```shell
  make errors
```
or entering your raspberry pi IP in the browser.
