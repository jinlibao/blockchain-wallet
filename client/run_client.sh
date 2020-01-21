#! /bin/sh
#
# run_client.sh
# Copyright (C) 2018 Libao Jin <jinlibao@outlook.com>
#
# Distributed under terms of the MIT license.
#

# store password for creating keyfiles

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
