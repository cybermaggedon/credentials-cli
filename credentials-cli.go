package main

/****************************************************************************

  Example credential management.  Interacts with Google APIs to
  download credentials.

  Usage:

    credentials-cli -A

      Authenticate with Google and download a Google OAUTH2 token to the
      .credentials-cli-auth file.  You should keep this token safe, and treat
      it as a password.

    credentials-cli -L

      Lists credentials you own.

    credentials-cli -D

      Downloads a credential.

****************************************************************************/

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	creds "github.com/cybermaggedon/credentials"
)

var options struct {
	Authenticate        bool   `short:"A" long:"authenticate" description:"Specifies to authenticate with Google API" required:"false"`
	ListCredentials     bool   `short:"L" long:"list-credentials" description:"Specifies to list credentials" required:"false"`
	TokenFile           string `short:"t" long:"token" description:"Name of token file" default:".credentials-cli-auth"`
	Bucket              string `short:"b" long:"bucket" description:"Name of credentials bucket" default:"example-credentials"`
	ListFormats         bool   `short:"F" long:"list-formats" description:"Specifies to list available formats for a credential" required:"false"`
	DownloadCredentials bool   `short:"D" long:"download-credentials" description:"Specifies to download credential" required:"false"`
	Project         string `short:"p" long:"project" description:"Project ID hosting credential services (CKMS and pub/sub)" default:"example"`
	Verbose             bool   `short:"v" long:"verbose" description:"Enable verbose output"`
	OutputFormat        string `short:"f" long:"output-format" description:"Output format, one of: pem, p12"`
	SvcAccountKey       string `short:"k" long:"svc-key" description:"Name of service accounts JSON file"`
	User                string `short:"u" long:"user" description:"Username"`
	Sign                bool   `short:"S" long:"sign" description:"Enable signing"`
	SigningCert         string `short:"C" long:"signing-cert" description:"Enable signing with signing cert"`
	SigningKey          string `short:"K" long:"signing-key" description:"Enable signing with signing key"`
	Identity            string `short:"i" long:"identity" description:"Identity - web/VPN name"`
	CreateVpn           bool   `long:"create-vpn" description:"Create VPN credential"`
	CreateWeb           bool   `long:"create-web" description:"Create web credential"`
	CreateProbe         bool   `long:"create-probe" description:"Create probe credential"`
	CreateVpnService    bool   `long:"create-vpn-service" description:"Create VPN service credential"`
	RevokeVpn           bool   `long:"revoke-vpn" description:"Revoke VPN credential"`
	RevokeWeb           bool   `long:"revoke-web" description:"Revoke web credential"`
	RevokeProbe         bool   `long:"revoke-probe" description:"Revoke probe credential"`
	RevokeVpnService    bool   `long:"revoke-vpn-service" description:"Revoke VPN service credential"`
	RevokeAll           bool   `long:"revoke-all" description:"Revoke all credentials"`
	McIdentifier        string `long:"mc-identifier" description:"mobileconfig identifier"`
	McName              string `long:"mc-name" description:"mobileconfig display name"`
	McDescription       string `long:"mc-description" description:"mobileconfig description"`
	Endpoint            string `long:"endpoint" description:"host:port for probe credential creation"`
	Hostname            string `long:"hostname" description:"host for VPN service credential creation"`
	Allocator           string `long:"allocator" description:"Address allocator host for VPN service credential creation"`
	Soc                 string `long:"soc" description:"SOC id for credential creation"`
}

// Fetches a user's account name.  This is the first email address listed
// in the EmailAddresses list in the profile.
func getMyEmail(client *creds.Client) (string, error) {

	if options.User != "" {
		return options.User, nil
	}

	return client.GetEmailAddress()

}

// CLI use-case: List credentials
func listCredentials() error {

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	// Query People API for my email address.
	email, err := getMyEmail(client)
	if err != nil {
		return err
	}

	// Get the creds index.
	credlist, err := client.GetIndex(email)
	if err != nil {
		return err
	}

	// Loop over creds, asking each to describe itself.
	for _, cred := range credlist {
		cred.(creds.Credential).Describe(os.Stdout, options.Verbose)
	}

	return nil

}

// CLI use-case: Download a cred.  Assumes user has already authorised
// and got a Google auth token.
func downloadCredential(credName string) error {

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	// What email address am I logged in with?
	email, err := getMyEmail(client)
	if err != nil {
		return err
	}

	// Fetch the credential index
	credlist, err := client.GetIndex(email)
	if err != nil {
		return err
	}

	// Will point to selected cred
	var selected creds.Credential

	// Loop through creds, searching for selected cred
	for _, cred := range credlist {
		if cred.(creds.Credential).GetId() == credName {
			selected = cred.(creds.Credential)
		}
	}

	// Error if no match
	if selected == nil {
		return errors.New("Credential not found")
	}

	// Download specified credential
	pays, err := selected.Get(client, options.OutputFormat)
	if err != nil {
		return err
	}

	for _, pay := range(pays) {
		if pay.Disposition == "store" {
			f, err := os.Create(pay.Filename)
			if err != nil {
				return err
			}
			f.Write(pay.Payload)
			f.Close()
			fmt.Printf("%s written to %s\n", pay.Description,
				pay.Filename)
		}
		if pay.Disposition == "show" {
			fmt.Printf("%s: %s\n", pay.Description,
				string(pay.Payload))
		}

	}

	return nil

}

// CLI use-case: Download a cred.  Assumes user has already authorised
// and got a Google auth token.
func listFormats(credName string) error {

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	// What email address am I logged in with?
	email, err := getMyEmail(client)
	if err != nil {
		return err
	}

	// Fetch the credential index
	credlist, err := client.GetIndex(email)
	if err != nil {
		return err
	}

	// Will point to selected cred
	var selected creds.Credential

	// Loop through creds, searching for selected cred
	for _, cred := range credlist {
		if cred.(creds.Credential).GetId() == credName {
			selected = cred.(creds.Credential)
		}
	}

	// Error if no match
	if selected == nil {
		return errors.New("Credential not found")
	}

	// Download specified credential
	fs := selected.GetFormats()
	if err != nil {
		return err
	}

	for _, f := range(fs) {
		fmt.Fprintf(os.Stdout, "%-20s - %s\n", f.Id, f.Description)
	}

	return nil

}

func getClient() (*creds.Client, error) {

	var client *creds.Client
	var err error
	
	if options.SvcAccountKey != "" {

		client, err = creds.NewSaClient(options.SvcAccountKey)
		if err != nil {
			return nil, err
		}
		
	} else {

		// Fetch the token from the offline storage.
		client, err = creds.NewClientFromTokenFile(options.TokenFile)
		if err != nil {
			return nil, err
		}

	}

	client.
		SetProject(options.Project).
		SetBucket(options.Bucket).
		SetMcIdentifier(options.McIdentifier).
		SetMcName(options.McName).
		SetMcDescription(options.McDescription)

	if options.Soc != "" {
		client.SetSoc(options.Soc)
	}

	if options.Sign {
		client.SetSigning(options.SigningKey, options.SigningCert)
	}

	if options.User != "" {
		client.SetUser(options.User)
	}

	return client, nil

}

func main() {

	// Parse arguments
	args, err := flags.ParseArgs(&options, os.Args)
	if err != nil {
		// ParseArgs already displays the error.  Don't need to
		// display it.
//		fmt.Println("Error:", err)
		os.Exit(1)
	}

	_ = args

	// CLI use-case: Authenticate
	if options.Authenticate {
		err = creds.Authenticate(options.TokenFile)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// CLI use-case: List credentials
	if options.ListCredentials {
		err = listCredentials()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// CLI use-case: Download credentials
	if options.DownloadCredentials {
		if len(args) < 2 {
			fmt.Println("Need to specify one or more credentials")
			os.Exit(1)
		}
		for _, cred := range args[1:] {
			err = downloadCredential(cred)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		}
		os.Exit(0)
	}

	// CLI use-case: Download credentials
	if options.ListFormats {
		if len(args) < 2 {
			fmt.Println("Need to specify one or more credentials")
			os.Exit(1)
		}
		for _, cred := range args[1:] {
			err = listFormats(cred)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		}
		os.Exit(0)
	}

	if options.CreateWeb {
		err := createWeb()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if options.CreateVpn {
		err = createVpn()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if options.CreateProbe {
		err = createProbe()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	
	if options.CreateVpnService {
		err = createVpnService()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if options.RevokeWeb {
		err = revokeWeb()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if options.RevokeVpn {
		err = revokeVpn()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if options.RevokeProbe {
		err = revokeProbe()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if options.RevokeVpnService {
		err = revokeVpnService()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if options.RevokeAll {
		err = revokeAll()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
	
}
