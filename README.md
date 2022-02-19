# VWAP Calculator
A real time [VWAP](https://en.wikipedia.org/wiki/Volume-weighted_average_price) calculator using the Coinbase websockets as data provider.
This calculate the VWAP by trading pair:

# Top level project structure
* `./coinbase/` contains all the code to connect to coinbase and get the real-time data. [docs](https://docs.cloud.coinbase.com/exchange/docs/websocket-overview)
* `./vwap/` contains all the logic to calculate the VWAP indicator using the data received from Coinbase.

# How to run the project
The command has two optionals flags
* `interval`: The sliding window to calculate the VWAP indicator. Default value: **200**
* `pairs`: a comma separated strings of cryptocurrencies pairs. Default value `BTC-USD,ETH-USD,ETH-BTC`

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
./vwap_calculator -interval=200 
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