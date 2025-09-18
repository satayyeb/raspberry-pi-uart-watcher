## Raspberry Pi UART watcher

![](https://images.ctfassets.net/3prze68gbwl1/asset-17suaysk1qa1hys/a23693cc21488367ff64bc1e7822d370/raspberry-pi-motion-sensor-detector.png)


**Overview:**

`raspberry-pi-uart-watcher` is a lightweight utility designed for Raspberry Pi systems to monitor and detect error logs originating from the UART interface. It continuously listens to UART communications and identifies any anomalies or errors, providing real-time alerts for prompt attention.

**Key Features:**

* **Real-Time Error Detection:** Monitors UART traffic and identifies error logs as they occur.
* **Web Interface:** Access error logs via a simple web interface by navigating to the Raspberry Pi's IP address in a browser.
* **Easy Deployment:** Deploy the application with a single command using the provided `Makefile`.

**Installation & Usage:**

1. Customize network environment variables in the `Makefile`.

2. Deploy the application by running:

   ```bash
   make deploy
   ```

3. View error logs by executing:

   ```bash
   make errors
   ```

   Alternatively, open your Raspberry Pi's IP address in a web browser to access the logs.

