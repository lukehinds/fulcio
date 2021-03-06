// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package config

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var validCfg = `
{
	"OIDCIssuers": {
		"foobar.com": {
			"IssuerURL": "foobar.com",
			"ClientID": "foo"
		}
	}
}
`

func TestLoad(t *testing.T) {
	td := t.TempDir()
	cfgPath := filepath.Join(td, "config.json")
	if err := ioutil.WriteFile(cfgPath, []byte(validCfg), 0644); err != nil {
		t.Fatal(err)
	}

	if err := Load(cfgPath); err != nil {
		t.Fatal(err)
	}

	cfg := Config()
	if got := cfg.OIDCIssuers["foobar.com"].ClientID; got != "foo" {
		t.Errorf("expected foo, got %s", got)
	}
	if got := cfg.OIDCIssuers["foobar.com"].IssuerURL; got != "foobar.com" {
		t.Errorf("expected foo, got %s", got)
	}
	if got := len(cfg.OIDCIssuers); got != 1 {
		t.Errorf("expected 1 issuer, got %d", got)
	}
}
func TestLoadDefaults(t *testing.T) {
	td := t.TempDir()

	// Don't put anything here!
	cfgPath := filepath.Join(td, "config.json")
	if err := Load(cfgPath); err != nil {
		t.Fatal(err)
	}

	cfg := Config()

	if diff := cmp.Diff(DefaultConfig, cfg); diff != "" {
		t.Errorf("DefaultConfig(): -want +got: %s", diff)

	}
}
