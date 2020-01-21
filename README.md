# Blockchain Wallet

## Overview

The blockchain keeps a "distributed ledger" (DLT) but has no information on who can update the ledger data.

This project has two parts: one is the server which executes commands to process transactions (creates transactions and writes it to the blockchain / retrieves transactions from the blockchain); and the other is the client which sends requests to the server to complete tasks such as checking the balance, sending funds to another account, listing accounts and etc.

## Set up

If this is your first time running go, you will need to set the `GOPATH`.

## Dependencies

At the root directory, run

```sh
go get ./
```

to download the dependencies. To make sure everything is properly set up, run

```
make test
```

at the root directory. If it passes all the tests, you are ready to go!

## To Start the Server

The server is a HTTP server. At the root directory, you can simply run the falling script to start the server:

```sh
make run_server
```

You will want to run this in it's own window as it runs until it is killed or until you send it a shutdown message.

## To Call the Server from the Client.

At the root directory, you can simply run the following script to start the client:

```sh
make run_client
```

Then go to `./client` folder, you can run the `./client` as follows:

```
Usage: ./client [ --cfg file ] [ --host URL ] [ --cmd Command ] [ --from Acct ] [ --to Acct ] [ --amount #### ] [ --addr Addr ] [ --password <PW> ]
Command can be:
  send-funds-to --from MyAcct --to AcctTo --amount ####
  list-accts
  list-wallet
  list-my-keys
  acct-value --acct Acct
  new-key-file --password <PW>
  acct-value --acct Acct
  shutdown-server --addr address --password <PW>
  validate-signed-message --addr address
  server-status
```

## Demo

After you run `make run_server`, the `./server/main/run_server.sh` would create a genesis block and some accounts with initial balances, and processes some transactions.

`./server/run_server`:
```sh
go build

account0=0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a
account1=0xb0d533a8064ed180967aa4dafa453deab107961c
account2=0x31568cc92115a2ebe6eb37e9a7c7f6334b988196
account3=0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300

rm ./data/*

# crate the genesis block
echo "\nCreate genesis block:"
./main --create-genesis

# test whether the block chain is successfully created
echo "\nCheck whether the genesis block has been created:"
./main --test-read-block

# list accounts available on the block chain
echo "\nList accounts on the block chain:"
./main --list-accounts

# check the balance for accounts before several transactions
echo "\nCheck the balances for some accounts:"
./main --show-balance $account0
./main --show-balance $account1
./main --show-balance $account2
./main --show-balance $account3

# test sending funds
echo "\nTransactions between some accounts:"
./main --test-send-funds $account0 $account1 500 x x x x
./main --test-send-funds $account0 $account2 500 x x x x
./main --test-send-funds $account0 $account3 500 x x x x

# check the balance for accounts after several transactions
echo "\nCheck the balances for the above accounts:"
./main --show-balance $account0
./main --show-balance $account1
./main --show-balance $account2
./main --show-balance $account3

# start the server at 127.0.0.1:9191
echo "\nStart the server at 127.0.0.1:9191:"
./main --server 127.0.0.1:9191
```

Here is the output:

```
( cd server/main ; ./run_server.sh)

Create genesis block:
((Mining)) Hash for Block [0543ce00776885535a69f621fd1b3a174d3a916c55b074d9414b15af0f3269b2] nonce [       0]
((Mining)) Hash for Block [0543ce00776885535a69f621fd1b3a174d3a916c55b074d9414b15af0f3269b2] nonce [       0]

Check whether the genesis block has been created:
You must run "create-genesis" first before this test.

List accounts on the block chain:

List of Addresses
	0x0c34a1a3c5ae302cb41f9cfd999e7950b8ebf40f
	0xe7b8a518bf1b5c4f01b2a7ee39a2800a982e06ee
	0x31568cc92115a2ebe6eb37e9a7c7f6334b988196
	0x5ae7b3cf64adc3d7fef099319a9be4acb8bd73ed
	0xb0d533a8064ed180967aa4dafa453deab107961c
	0x42f487a6d5c86962310d5ab5afe5cad7bc80805b
	0x40681739b0ef568acce20f5575ad4cf24223926f
	0x6e06bf940bb57ade69cb03153d1c3842411bd3c1
	0x885765a2fcfb72e68d82d656c6411b7d27bacdd7
	0x4af64cd87a47aab7cffdbada6bfd6aef47036c03
	0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300
	0xdb180da9a8982c7bb75ca40039f959cb959c62e8
	0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a
	0x9d41e5938767466af28865e1c33071f1561d57a8
	0x3b65b88e4256c8926358551072f17460efe5452b

Check the balances for some accounts:
Acct: 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a Value: 500000
Acct: 0xb0d533a8064ed180967aa4dafa453deab107961c Value: 500000
Acct: 0x31568cc92115a2ebe6eb37e9a7c7f6334b988196 Value: 500000
Acct: 0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300 Value: 500000

Transactions between some accounts:

Before sending funds the balances for 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a and 0xb0d533a8064ed180967aa4dafa453deab107961c are as follows:
Acct: 0xb0d533a8064ed180967aa4dafa453deab107961c Value: 500000
Acct: 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a Value: 500000
Sent 0xb0d533a8064ed180967aa4dafa453deab107961c $500 from 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a
((Mining)) Hash for Block [001a443254bbf24ca1cb60f06f909a40cc01ea6942dbce0a474b241f9876d3d6] nonce [       3]

Before sending funds the balances for 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a and 0x31568cc92115a2ebe6eb37e9a7c7f6334b988196 are as follows:
Acct: 0x31568cc92115a2ebe6eb37e9a7c7f6334b988196 Value: 500000
Acct: 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a Value: 499500
Sent 0x31568cc92115a2ebe6eb37e9a7c7f6334b988196 $500 from 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a
((Mining)) Hash for Block [0fb66d13d109ab4cbd87f184e30915ad43e65e0423383d559a669d88af550122] nonce [      11]

Before sending funds the balances for 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a and 0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300 are as follows:
Acct: 0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300 Value: 500000
Acct: 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a Value: 499000
Sent 0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300 $500 from 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a
((Mining)) Hash for Block [043b518a9a5ae00b872040df7d22848769e207a48ca3a261a001849379e3e138] nonce [      41]

Check the balances for the above accounts:
Acct: 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a Value: 498500
Acct: 0xb0d533a8064ed180967aa4dafa453deab107961c Value: 500500
Acct: 0x31568cc92115a2ebe6eb37e9a7c7f6334b988196 Value: 500500
Acct: 0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300 Value: 500500
```

Then it would start the server at `http://localhost:9191` and listen to the client. Next, running `make run_client` is actually calling `./client/run_client.sh`, which creates some transactions or requests to the server.

`run_client.sh`:

```sh
go build

password=123456
loginAccount=0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a
account1=0xb0d533a8064ed180967aa4dafa453deab107961c
account2=0x31568cc92115a2ebe6eb37e9a7c7f6334b988196
account3=0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300

if [[ -d wallet-data ]]; then
    rm -rf wallet-data
else
    make -p wallet-data
fi

# test call
./client --cmd echo

# list accounts available on server
echo "\nList accounts on the server:"
./client --cmd list-accts

# check the server status
echo "\nCheck server status:"
./client --cmd server-status

# generate 1 key pair and store it
echo "\nGenerate one key pair and store it:"
./client --cmd new-key-file --password $password
addressOfKeyfile=$(./client --cmd list-my-keys)

# check the balance of several accounts
echo "\nCheck balance:"
./client --cmd acct-value --acct $loginAccount
./client --cmd acct-value --acct $account1
./client --cmd acct-value --acct $account2
./client --cmd acct-value --acct $account3

# validate the signed message
echo "\nValidate the signed message:"
./client --cmd validate-signed-message --addr $addressOfKeyfile --password $password

# test of send funds to
echo "\nSend \$1000 from $loginAccount to $account1:"
./client --cmd send-funds-to --from $loginAccount --to $account1 --amount 1000 --memo dinner --addr $addressOfKeyfile --password $password
echo "\nSend \$1500 from $loginAccount to $account2:"
./client --cmd send-funds-to --from $loginAccount --to $account2 --amount 1500 --memo dinner --addr $addressOfKeyfile --password $password
echo "\nSend \$2500 from $loginAccount to $account3:"
./client --cmd send-funds-to --from $loginAccount --to $account3 --amount 2500 --memo dinner --addr $addressOfKeyfile --password $password

# check the balance of accounts
echo "\nCheck balance:"
./client --cmd acct-value --acct $loginAccount
./client --cmd acct-value --acct $account1
./client --cmd acct-value --acct $account2
./client --cmd acct-value --acct $account3

# shut down the server
./client --cmd shutdown-server --addr $addressOfKeyfile --password $password
```

### The Output on the Server Side

```
Before sending funds the balances for 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a and 0xb0d533a8064ed180967aa4dafa453deab107961c are as follows:
Acct: 0xb0d533a8064ed180967aa4dafa453deab107961c Value: 500500
Acct: 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a Value: 498500
Sent 0xb0d533a8064ed180967aa4dafa453deab107961c $1000 from 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a
((Mining)) Hash for Block [0fcbcae92e1dbc85576d9d38ac8816a6904c6b508d13e688c696bebd5492a74f] nonce [       0]

Before sending funds the balances for 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a and 0x31568cc92115a2ebe6eb37e9a7c7f6334b988196 are as follows:
Acct: 0x31568cc92115a2ebe6eb37e9a7c7f6334b988196 Value: 500500
Acct: 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a Value: 497500
Sent 0x31568cc92115a2ebe6eb37e9a7c7f6334b988196 $1500 from 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a
((Mining)) Hash for Block [0f4d1bd72a6f8797c151c9071105f701717639b582401b2e91d208a6c9f509da] nonce [      13]

Before sending funds the balances for 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a and 0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300 are as follows:
Acct: 0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300 Value: 500500
Acct: 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a Value: 496000
Sent 0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300 $2500 from 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a
((Mining)) Hash for Block [099919cb7ab58d459c86ab7a61dfe739258bb8e35b1c275703c0b8c18a643cf1] nonce [      18]
Shutdown Now
```

### The Output on the Client Side

```
( cd client ; ./run_client.sh)
Echo was called

List accounts on the server:
Body: [
	"0xb0d533a8064ed180967aa4dafa453deab107961c",
	"0x3b65b88e4256c8926358551072f17460efe5452b",
	"0x31568cc92115a2ebe6eb37e9a7c7f6334b988196",
	"0xdb180da9a8982c7bb75ca40039f959cb959c62e8",
	"0x5ae7b3cf64adc3d7fef099319a9be4acb8bd73ed",
	"0x885765a2fcfb72e68d82d656c6411b7d27bacdd7",
	"0x9d41e5938767466af28865e1c33071f1561d57a8",
	"0x40681739b0ef568acce20f5575ad4cf24223926f",
	"0x4af64cd87a47aab7cffdbada6bfd6aef47036c03",
	"0xe7b8a518bf1b5c4f01b2a7ee39a2800a982e06ee",
	"0x6e06bf940bb57ade69cb03153d1c3842411bd3c1",
	"0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a",
	"0x0c34a1a3c5ae302cb41f9cfd999e7950b8ebf40f",
	"0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300",
	"0x42f487a6d5c86962310d5ab5afe5cad7bc80805b"
]

Check server status:
Body: {"status":"success","name":"go-server version 1.0.0","URI":"/api/status","req":Error:json: unsupported type: func() (io.ReadCloser, error), "response_header":{
	"Content-Type": [
		"application/json"
	]
}}

Generate one key pair and store it:
Address: 0x79fc57A68387aCeB79f0bb610F751eDFE04bF7C2
File Name: wallet-data/UTC--2020-01-21T03-21-11.338725Z--79fc57A68387aCeB79f0bb610F751eDFE04bF7C2

Check balance:
Body: { "status":"success", "acct": "0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a"
  "value": 498500 }

Body: { "status":"success", "acct": "0xb0d533a8064ed180967aa4dafa453deab107961c"
  "value": 500500 }

Body: { "status":"success", "acct": "0x31568cc92115a2ebe6eb37e9a7c7f6334b988196"
  "value": 500500 }

Body: { "status":"success", "acct": "0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300"
  "value": 500500 }


Validate the signed message:
Match of Addr [79fc57A68387aCeB79f0bb610F751eDFE04bF7C2] to fn [UTC--2020-01-21T03-21-11.338725Z--79fc57A68387aCeB79f0bb610F751eDFE04bF7C2]
Body: {"status":"success","msg":"Signature validated"}


Send $1000 from 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a to 0xb0d533a8064ed180967aa4dafa453deab107961c:
Match of Addr [79fc57A68387aCeB79f0bb610F751eDFE04bF7C2] to fn [UTC--2020-01-21T03-21-11.338725Z--79fc57A68387aCeB79f0bb610F751eDFE04bF7C2]
http://0x7e3aFEc048bC7be745d0fA0F5af97D3978C40E9A:237801@127.0.0.1:9191/api/send-funds-to
Body: {"status":"success", "blockNo":4 }

Send $1500 from 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a to 0x31568cc92115a2ebe6eb37e9a7c7f6334b988196:
Match of Addr [79fc57A68387aCeB79f0bb610F751eDFE04bF7C2] to fn [UTC--2020-01-21T03-21-11.338725Z--79fc57A68387aCeB79f0bb610F751eDFE04bF7C2]
http://0x7e3aFEc048bC7be745d0fA0F5af97D3978C40E9A:237801@127.0.0.1:9191/api/send-funds-to
Body: {"status":"success", "blockNo":5 }

Send $2500 from 0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a to 0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300:
Match of Addr [79fc57A68387aCeB79f0bb610F751eDFE04bF7C2] to fn [UTC--2020-01-21T03-21-11.338725Z--79fc57A68387aCeB79f0bb610F751eDFE04bF7C2]
http://0x7e3aFEc048bC7be745d0fA0F5af97D3978C40E9A:237801@127.0.0.1:9191/api/send-funds-to
Body: {"status":"success", "blockNo":6 }

Check balance:
Body: { "status":"success", "acct": "0x7e3afec048bc7be745d0fa0f5af97d3978c40e9a"
  "value": 493500 }

Body: { "status":"success", "acct": "0xb0d533a8064ed180967aa4dafa453deab107961c"
  "value": 501500 }

Body: { "status":"success", "acct": "0x31568cc92115a2ebe6eb37e9a7c7f6334b988196"
  "value": 502000 }

Body: { "status":"success", "acct": "0x4e27d9c8ba3a6904f7a7cb31eae5ccce8bf33300"
  "value": 503000 }

Match of Addr [79fc57A68387aCeB79f0bb610F751eDFE04bF7C2] to fn [UTC--2020-01-21T03-21-11.338725Z--79fc57A68387aCeB79f0bb610F751eDFE04bF7C2]
Error: 500
```
