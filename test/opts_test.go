// Copyright 2015-2016 Apcera Inc. All rights reserved.

package test

import (
	"testing"
	"crypto/tls"
)

func TestServerConfig(t *testing.T) {
	srv, opts := RunServerWithConfig("./configs/override.conf")
	defer srv.Shutdown()

	c := createClientConn(t, opts.Host, opts.Port)
	defer c.Close()

	sinfo := checkInfoMsg(t, c)
	if sinfo.MaxPayload != opts.MaxPayload {
		t.Fatalf("Expected max_payload from server, got %d vs %d",
			opts.MaxPayload, sinfo.MaxPayload)
	}
}

func TestTLSConfig(t *testing.T) {
	srv, opts := RunServerWithConfig("./configs/tls.conf")
	defer srv.Shutdown()

	c := createClientConn(t, opts.Host, opts.Port)
	defer c.Close()

	tlsc := tls.Client(c, &tls.Config{InsecureSkipVerify: true})
	err := tlsc.Handshake()
	if err != nil {
		t.Fatal(err)
	}
	sinfo := checkInfoMsg(t, tlsc)
	if !sinfo.TLSRequired {
		t.Fatal("Expected TLSRequired to be true when configured")
	}
}
