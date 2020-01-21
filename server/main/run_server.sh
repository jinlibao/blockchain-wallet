#! /bin/sh
#
# run.sh
# Copyright (C) 2018 Libao Jin <jinlibao@outlook.com>
#
# Distributed under terms of the MIT license.
#

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
