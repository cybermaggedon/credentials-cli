/****************************************************************************

  Management: creation and revocation of creds

****************************************************************************/

package main

import (
	"errors"
//	"fmt"
//	"google.golang.org/api/pubsub/v1"
)

// Web credential creation
func createWeb() error {

	if options.User == "" {
		return errors.New("Must specify user")
	}
	if options.Identity == "" {
		return errors.New("Must specify identity")
	}

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	err = client.CreateWebCredential(options.User, options.Identity)
	if err != nil {
		return err
	}

	return nil

}

// Web credential creation
func createVpn() error {

	if options.User == "" {
		return errors.New("Must specify user")
	}
	if options.Identity == "" {
		return errors.New("Must specify identity")
	}

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	err = client.CreateVpnCredential(options.User, options.Identity)
	if err != nil {
		return err
	}

	return nil

}
// Web credential creation
func createVpnService() error {

	if options.User == "" {
		return errors.New("Must specify user")
	}
	if options.Identity == "" {
		return errors.New("Must specify identity")
	}
	if options.Allocator == "" {
		return errors.New("Must specify allocator")
	}
	if options.Hostname == "" {
		return errors.New("Must specify hostname")
	}

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	err = client.CreateVpnServiceCredential(options.User, options.Identity,
		options.Hostname, options.Allocator)
	if err != nil {
		return err
	}

	return nil

}

// Web credential creation
func createProbe() error {

	if options.User == "" {
		return errors.New("Must specify user")
	}
	if options.Identity == "" {
		return errors.New("Must specify identity")
	}
	if options.Endpoint == "" {
		return errors.New("Must specify endpoint")
	}

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	err = client.CreateProbeCredential(options.User, options.Identity,
		options.Endpoint)
	if err != nil {
		return err
	}

	return nil

}

// Web credential revocation
func revokeWeb() error {

	if options.User == "" {
		return errors.New("Must specify user")
	}
	if options.Identity == "" {
		return errors.New("Must specify identity")
	}

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	err = client.RevokeWebCredential(options.User, options.Identity)
	if err != nil {
		return err
	}
	
	return nil
}

// Web credential revocation
func revokeVpn() error {

	if options.User == "" {
		return errors.New("Must specify user")
	}
	if options.Identity == "" {
		return errors.New("Must specify identity")
	}

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	err = client.RevokeVpnCredential(options.User, options.Identity)
	if err != nil {
		return err
	}
	
	return nil
}

// Probe credential revocation
func revokeProbe() error {

	if options.User == "" {
		return errors.New("Must specify user")
	}
	if options.Identity == "" {
		return errors.New("Must specify identity")
	}

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	err = client.RevokeProbeCredential(options.User, options.Identity)
	if err != nil {
		return err
	}
	
	return nil
}

// Web credential revocation
func revokeVpnService() error {

	if options.User == "" {
		return errors.New("Must specify user")
	}
	if options.Identity == "" {
		return errors.New("Must specify identity")
	}

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	err = client.RevokeVpnServiceCredential(options.User, options.Identity)
	if err != nil {
		return err
	}
	
	return nil
}

// Revoke all
func revokeAll() error {

	if options.User == "" {
		return errors.New("Must specify user")
	}
	if options.Identity == "" {
		return errors.New("Must specify identity")
	}

	// Get Google API client
	client, err := getClient()
	if err != nil {
		return err
	}

	err = client.RevokeAll(options.User)
	if err != nil {
		return err
	}
	
	return nil

}
