# tw-ethereum
Query interface for ethereum to get list of transactions for single user

# Description
Query interface for ethereum that follows the parser interface provided by TW

### Parser interface
```
type Parser interface {
	// last parsed block
	GetCurrentBlock() int
	// add address to observer
	Subscribe(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []Transaction
}
```

With the current implementation, I have changed the interface to:
```
type Parser interface {
	// last parsed block
	GetCurrentBlock() (int, error)
	// add address to observer
	Subscribe(address string) (bool, error)
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) ([]Transaction, error)
}
```
This allows the propagation of errors from the underlying back up


### Usage
The build system used in this application is bazel.  To run build or test or run, one will need to
run the following to install `bazelisk` which will install a copy of bazel at runtime.

MacOSX
```
# Assuming user has installed brew from https://brew.sh
brew install bazelisk
```

Once bazelisk is installed, you will need to run the program in the following ways:

#### Main with help
```
bazel run //cmd/ethtx
bazel run //cmd/ethtx help
```

The cobra cli also provides shell completion scripts that help with 

#### Get Current block
```
bazel run //cmd/ethtx block
```

#### Get List of Transactions - Unsubscribed contract/address
```
bazel run //cmd/ethtx list --addr "<address>"
```

#### Add an address to monitor
```
bazel run //cmd/ethtx subscribe --addr "<address>"
```
This should run the initial getfilterlog, and save the subscription list somewhere

#### Get latest list of changes for transactions
```
bazel run //cmd/ethtx subscribe --addr "<address>"
bazel run //cmd/ethtx list --addr "<address>"
```
