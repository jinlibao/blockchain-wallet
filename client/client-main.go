package main

// Client Program
//
// Options
//   --host http://127.0.0.1:9191
//   --cmd send-funds-to --from MyAcct --to AcctTo --amount ####
//   --cmd list-accts
//   --cmd list-wallet
//   --cmd list-my-keys
//   --cmd acct-value --acct Acct
//   --cmd new-key-file --password <PW>
//   --cmd shutdown-server --addr address --password <PW>
//   --cmd validate-signed-message --addr address
//   --cmd server-status
//
//   --help
//

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

type ConfigData struct {
	Host       string
	WalletPath string
	LoginAcct  string
	LoginPin   string
}

var GCfg ConfigData

var Cfg = flag.String("cfg", "cfg.json", "config file for this program.")
var Host = flag.String("host", "", "server host - ip address:port")
var Cmd = flag.String("cmd", "", "command to run.")
var From = flag.String("from", "", "from account.")
var To = flag.String("to", "", "to account.")
var Acct = flag.String("acct", "", "account to specify.")
var Addr = flag.String("addr", "", "address to the keyfile")
var Amount = flag.String("amount", "", "amount of money to use in tranaction.")
var Password = flag.String("password", "", "password to use if creating a key file.")
var Memo = flag.String("memo", "", "Memo for send funds tranaction.")

var HostWithUnPw string

func main() {
	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()
	if len(fns) > 0 {
		usage()
	}

	GCfg = ReadCfg(*Cfg)

	os.MkdirAll(GCfg.WalletPath, 0755)

	uP, err := url.Parse(GCfg.Host)
	if len(*Host) > 0 {
		uP, err = url.Parse(*Host)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse the host, error:%s\n", err)
		os.Exit(1)
	}
	uP.User = url.UserPassword(GCfg.LoginAcct, GCfg.LoginPin) // Note RFC 2396 - this is very bad security!
	HostWithUnPw = fmt.Sprintf("%s", uP)
	// fmt.Printf("HostWithUnPw= ->%s<-\n", HostWithUnPw)

	switch *Cmd {
	case "echo":
		fmt.Printf("Echo was called\n")

	case "list-accts":
		urlStr := fmt.Sprintf("%s/api/acct-list", HostWithUnPw)
		// fmt.Printf("urlStr= ->%s<-\n", urlStr)
		status, body := DoGet(urlStr)
		if status == 200 {
			fmt.Printf("Body: %s\n", body)
		} else {
			fmt.Printf("Error: %d\n", status)
		}

	case "shutdown-server":
		RequiredOption("addr", *Addr)
		RequiredOption("password", *Password)
		password := getPassphrase(*Password)
		keyFile := getKeyFileFromAddr(*Addr)
		data, err := ioutil.ReadFile(keyFile)
		if err != nil {
			fmt.Errorf("unable to read keyfile %s Error:%s", keyFile, err)
		}
		key, err := keystore.DecryptKey(data, password)
		if err != nil {
			fmt.Errorf("unable to decrypt %s Error:%s", keyFile, err)
		}
		inMessage, err := GenRandBytes(20)
		if err != nil {
			fmt.Errorf("unable to generate random message Error:%s", err)
		}
		message := hex.EncodeToString(inMessage)
		rawSignature, err := crypto.Sign(signHash(inMessage), key.PrivateKey)
		if err != nil {
			fmt.Errorf("unable to sign message Error:%s", err)
		}
		signature := hex.EncodeToString(rawSignature)

		urlStr := fmt.Sprintf("%s/api/shutdown", HostWithUnPw)
		status, body := DoGet(urlStr, "addr", *Addr, "signature", signature, "msg", message)
		// status, body := DoGet(urlStr)
		if status == 200 {
			fmt.Printf("Body: %s\n", body)
		} else {
			fmt.Printf("Error: %d\n", status)
		}

	case "server-status":
		urlStr := fmt.Sprintf("%s/api/status", HostWithUnPw)
		status, body := DoGet(urlStr)
		if status == 200 {
			fmt.Printf("Body: %s\n", body)
		} else {
			fmt.Printf("Error: %d\n", status)
		}

	case "acct-value":
		RequiredOption("acct", *Acct)
		urlStr := fmt.Sprintf("%s/api/acct-value", HostWithUnPw)
		// fmt.Printf("Client-AT: %s acct ->%s<-\n", godebug.LF(), *Acct)
		status, body := DoGet(urlStr, "acct", *Acct)
		if status == 200 {
			fmt.Printf("Body: %s\n", body)
		} else {
			fmt.Printf("Error: %d\n", status)
		}

	case "new-key-file":
		password := getPassphrase(*Password)
		if err := GenerateNewKeyFile(password); err != nil {
			fmt.Fprintf(os.Stderr, "Error generating KeyFile, Error:%s\n", err)
			os.Exit(1)
		}

	case "list-my-keys", "list-wallet":
		fns, _ := GetFilenames(GCfg.WalletPath)
		for _, fn := range fns {
			// fmt.Printf("%s\n", fn) // TODO - clean up file name into just the "account" part.
			words := strings.Split(fn, "-")
			fmt.Printf("0x%s\n", words[len(words)-1])
		}

	case "validate-signed-message": // call the server with a signed message.  Verify if the message is properly signed.
		// Replace the call below with your code - call your own function.
		// InstructorValidateSignedMessage(*Acct, *Password)
		RequiredOption("addr", *Addr)
		RequiredOption("password", *Password)

		password := getPassphrase(*Password)
		keyFile := getKeyFileFromAddr(*Addr)
		data, err := ioutil.ReadFile(keyFile)
		if err != nil {
			fmt.Errorf("unable to read keyfile %s Error:%s", keyFile, err)
		}
		key, err := keystore.DecryptKey(data, password)
		if err != nil {
			fmt.Errorf("unable to decrypt %s Error:%s", keyFile, err)
		}
		inMessage, err := GenRandBytes(20)
		if err != nil {
			fmt.Errorf("unable to generate random message Error:%s", err)
		}
		message := hex.EncodeToString(inMessage)
		rawSignature, err := crypto.Sign(signHash(inMessage), key.PrivateKey)
		if err != nil {
			fmt.Errorf("unable to sign message Error:%s", err)
		}
		signature := hex.EncodeToString(rawSignature)

		urlStr := fmt.Sprintf("%s/api/validate-signed-message", HostWithUnPw)
		status, body := DoGet(urlStr, "addr", *Addr, "signature", signature, "msg", message)
		if status == 200 {
			fmt.Printf("Body: %s\n", body)
		} else {
			fmt.Printf("Error: %d\n", status)
		}

	case "send-funds-to":
		// Replace the call below with your code - call your own function.
		// InstructorSendFundsTo(*From, *To, *Password, *Memo, *Amount)
		RequiredOption("from", *From)
		RequiredOption("to", *To)
		RequiredOption("amount", *Amount)
		RequiredOption("memo", *Memo)
		RequiredOption("addr", *Addr)
		RequiredOption("password", *Password)

		keyFile := getKeyFileFromAddr(*Addr)
		password := getPassphrase(*Password)
		data, err := ioutil.ReadFile(keyFile)
		if err != nil {
			fmt.Errorf("unable to read keyfile %s Error:%s", keyFile, err)
		}
		key, err := keystore.DecryptKey(data, password)
		if err != nil {
			fmt.Errorf("unable to decrypt %s Error:%s", keyFile, err)
		}

		urlStr := fmt.Sprintf("%s/api/send-funds-to", HostWithUnPw)
		fmt.Println(urlStr)
		inMessage := crypto.Keccak256([]byte(urlStr))
		message := hex.EncodeToString(inMessage)
		rawSignature, err := crypto.Sign(signHash(inMessage), key.PrivateKey)
		if err != nil {
			fmt.Errorf("unable to sign message Error:%s", err)
		}
		signature := hex.EncodeToString(rawSignature)

		status, body := DoGet(urlStr, "from", *From, "to", *To, "amount", *Amount, "addr", *Addr, "signature", signature, "msg", message, "memo", *Memo)
		if status == 200 {
			fmt.Printf("Body: %s\n", body)
		} else {
			fmt.Printf("Error: %d\n", status)
		}

	default:
		usage()
	}
}

func getKeyFileFromAddr(addr string) (keyFile string) {

	if addr[0:2] == "0x" || addr[0:2] == "0X" {
		addr = addr[2:]
	}

	fns, _ := GetFilenames(GCfg.WalletPath) // List of Files, discard any directories.
	for _, fn := range fns {
		if MatchAddrToFilename(addr, fn) {
			fmt.Printf("Match of Addr [%s] to fn [%s]\n", addr, fn)
			return filepath.Join(GCfg.WalletPath, fn)
		}
	}

	return
}

func MatchAddrToFilename(addr, fn string) bool {
	re, err := regexp.Compile(fmt.Sprintf("(?i)%s", addr)) // compare, ignore case.
	if err != nil {
		fmt.Printf("Unable to process matching of account to file name, addr [%s], fn [%s] error [%s]\n", addr, fn, err)
		os.Exit(1)
	}
	return re.MatchString(fn)
}

func RequiredOption(name, value string) {
	if value == "" {
		fmt.Fprintf(os.Stderr, "%s is a required option\n", name)
		os.Exit(1)
	}
}

func RequiredOptionInt(name string, value int) {
	if value <= 0 {
		fmt.Fprintf(os.Stderr, "%s is a required option\n", name)
		os.Exit(1)
	}
}

func ReadCfg(fn string) (rv ConfigData) {
	// Set defaults.
	rv.Host = "http://127.0.0.1:9191"
	rv.WalletPath = "./wallet-data"

	buf, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read %s Error:%s\n", fn, err)
		os.Exit(1)
	}
	err = json.Unmarshal(buf, &rv)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid initialization - Unable to parse JSON file, %s\n", err)
		os.Exit(1)
	}
	return
}

// usage will print a usage message and exit.
func usage() {
	fmt.Printf("Usage: client [ --cfg file ] [ --host URL ] [ --cmd Command ] [ --from Acct ] [ --to Acct ] [ --amount #### ] [ --addr Addr ] [ --password <PW> ]\n")
	fmt.Printf("Command can be:\n")
	fmt.Printf("  send-funds-to --from MyAcct --to AcctTo --amount ####\n")
	fmt.Printf("  list-accts\n")
	fmt.Printf("  list-wallet\n")
	fmt.Printf("  list-my-keys\n")
	fmt.Printf("  acct-value --acct Acct\n")
	fmt.Printf("  new-key-file --password <PW>\n")
	fmt.Printf("  acct-value --acct Acct\n")
	fmt.Printf("  shutdown-server --addr address --password <PW>\n")
	fmt.Printf("  validate-signed-message --addr address\n")
	fmt.Printf("  server-status\n")
	os.Exit(1)
}

// DoGet performs a GET operation over HTTP.
func DoGet(uri string, args ...string) (status int, rv string) {

	sep := "?"
	var qq bytes.Buffer
	qq.WriteString(uri)
	for ii := 0; ii < len(args); ii += 2 {
		// q = q + sep + name + "=" + value;
		qq.WriteString(sep)
		qq.WriteString(url.QueryEscape(args[ii]))
		qq.WriteString("=")
		if ii < len(args) {
			qq.WriteString(url.QueryEscape(args[ii+1]))
		}
		sep = "&"
	}
	url_q := qq.String()

	// fmt.Printf("Client-AT: %s, url=%s\n", godebug.LF(), url_q)

	res, err := http.Get(url_q)
	if err != nil {
		return 500, ""
	} else {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Error returnd: %s\n", err)
			return 500, ""
		}
		status = res.StatusCode
		if status == 200 {
			rv = string(body)
		}
		return
	}
}

// GenRandBytes will generate nRandBytes of random data using the random reader.
func GenRandBytes(nRandBytes int) (buf []byte, err error) {
	buf = make([]byte, nRandBytes)
	_, err = rand.Read(buf)
	if err != nil {
		return nil, err
	}
	return
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetFilenames(dir string) (filenames, dirs []string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, nil
	}
	for _, fstat := range files {
		if !strings.HasPrefix(string(fstat.Name()), ".") {
			if fstat.IsDir() {
				dirs = append(dirs, fstat.Name())
			} else {
				filenames = append(filenames, fstat.Name())
			}
		}
	}
	return
}
