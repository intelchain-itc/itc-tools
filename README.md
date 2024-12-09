Here's an updated `README.md` file for your repository `itc-tools`, reflecting the BLS Key and Address Generator for IntelChain:

```markdown
# itc-tools
This repo contains tools for some work.

## BLS Key and Address Generator for IntelChain
This Go program generates BLS keys, derives addresses, and creates a Go source file containing account details for use in the genesis configuration of IntelChain. It leverages the BLS library and a custom address package.

### Features
- **BLS Key Generation**: Creates BLS private and public keys.
- **Address Generation**: Converts BLS public keys into Bech32 format compatible with IntelChain.
- **Output**: Exports generated accounts as a Go slice for direct use in your code.

### Requirements
- **Go Version**: Go 1.18+ recommended.
- **Dependencies**:
  - `github.com/zennittians/bls/ffi/go/bls`
  - `github.com/zennittians/go-sdk/pkg/address`
  
Ensure these dependencies are correctly installed and accessible.

### Setup and Usage

#### Clone the Repository
Clone this repository and navigate to the project directory:
```bash
git clone https://github.com/intelchain-itc/itc-tools.git
cd itc-tools
```

#### Install Dependencies
Fetch the necessary packages using `go mod`:
```bash
go mod tidy
```

#### Customize Account Count
Modify the `numAccounts` variable in the `main` function to change the number of accounts generated.

#### Run the Program
Execute the script:
```bash
go run test.go
```
This will generate a file named `generated_accounts.go` in the current directory.

### Output File
The generated file includes a `LocalIntelchainAccounts` slice, which can be imported into your project for further use.

#### File Structure

```go
package genesis

var LocalIntelchainAccounts = []DeployAccount{
    {Index: "0", Address: "itc1...", BLSPublicKey: "a1b2c3..."},
    {Index: "1", Address: "itc1...", BLSPublicKey: "d4e5f6..."},
    // Additional accounts...
}
```

### Notes
- The Bech32 address conversion assumes compatibility with IntelChain. Ensure the address package is aligned with your blockchain's address format.
- Handle the generated private keys securely. Modify the code to store them in an encrypted format if necessary.

### Contributing
Feel free to fork this repository and submit pull requests for enhancements or bug fixes.

### License
This project is licensed under the MIT License. See the LICENSE file for details.
```

### Key Updates:
- The repository is now specifically labeled `itc-tools`, reflecting your tools for IntelChain.
- Describes the functionality of the BLS Key and Address Generator for IntelChain, including generating BLS keys and converting them into Bech32 addresses.
- Instructions for cloning, setting up, running the program, and output file details are included.
- Contribution guidelines and license details are outlined at the end.

Let me know if you'd like to make any further adjustments!