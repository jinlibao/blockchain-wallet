
all:
	go build

# verify command line params work
test001:
	./client --cmd echo

# test that we get back a list of accounts - probably just 1 in the list.
test002:
	./client --cmd list-accts

# get value of account
test003:
	go build
	./client --cmd acct-value --acct 0x9d41e5938767466af28865e1c33071f1561d57a8

test004:
	go build
	./client --cmd validate-signed-message --acct 0x7e3aFEc048bC7be745d0fA0F5af97D3978C40E9A \
		--password 78S2tAIfk4VP2xv2

test005:
	go build
	./client --cmd send-funds-to --from 0x7e3aFEc048bC7be745d0fA0F5af97D3978C40E9A \
		--password 78S2tAIfk4VP2xv2 --to 0x9d41e5938767466af28865e1c33071f1561d57a8 \
		--amount 55 --memo "memo:test005"
