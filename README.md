# VWAP Calculator
A real time [VWAP](https://en.wikipedia.org/wiki/Volume-weighted_average_price) calculator using the Coinbase websockets as data provider. This calculate the VWAP by trading pair.

An output example:
```bash
vwap.calculator | 2022/02/19 20:20:07 websocket connected to: wss://ws-feed.exchange.coinbase.com/
vwap.calculator | 2022/02/19 20:20:07 collecting datapoints for pairs: BTC-USD,ETH-USD,ETH-BTC | interval: 200
vwap.calculator | 2022/02/19 20:20:08 map[BTC-USD:40065.189999999995 ETH-BTC:0.06886 ETH-USD:2758.5899999999997]
vwap.calculator | 2022/02/19 20:20:08 map[BTC-USD:40065.189999999995 ETH-BTC:0.06886 ETH-USD:2758.5899999999997]
vwap.calculator | 2022/02/19 20:20:08 map[BTC-USD:40065.189999999995 ETH-BTC:0.06886 ETH-USD:2758.59]
vwap.calculator | 2022/02/19 20:20:09 map[BTC-USD:40065.19 ETH-BTC:0.06886 ETH-USD:2758.59]
vwap.calculator | 2022/02/19 20:20:09 map[BTC-USD:40065.19 ETH-BTC:0.06886 ETH-USD:2758.59]
vwap.calculator | 2022/02/19 20:20:09 map[BTC-USD:40065.189999999995 ETH-BTC:0.06886 ETH-USD:2758.59]
vwap.calculator | 2022/02/19 20:20:09 map[BTC-USD:40065.19520585874 ETH-BTC:0.06886 ETH-USD:2758.59]
vwap.calculator | 2022/02/19 20:20:09 map[BTC-USD:40065.1960835862 ETH-BTC:0.06886 ETH-USD:2758.59]
...
```

# Top level project structure
* `./coinbase/` contains all the code to connect to coinbase and get the real-time data. [docs](https://docs.cloud.coinbase.com/exchange/docs/websocket-overview)
* `./vwap/` contains all the logic to calculate the VWAP indicator using the data received from Coinbase.

# How to run the project
The command has two optionals flags
* `interval`: The sliding window to calculate the VWAP indicator. Default value: **200**
* `pairs`: a comma separated strings of cryptocurrencies pairs. Default value `BTC-USD,ETH-USD,ETH-BTC`

```bash
./vwap_calculator -interval=50 -pairs=SOL-USD
```

## Using Go 
make sure that you have go version 1.17
```
make run
```

## Compiling first

```bash
make build
``` 

and then

```bash
./vwap_calculator
```

and finally 

```bash
make clean
``` 

## Using Docker Compose
The VWAP calculator will run automatically

```bash
docker-compose up --build
``` 

# How to run the tests
You can execute the entire suit running the following command

```bash
make test
```

# Other comments
* This project run the tests using [GitHub actions](https://github.com/features/actions). Check the workflow runs [here](https://github.com/jedi4z/vwap-calculator/actions)