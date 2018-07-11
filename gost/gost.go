/* Vuls - Vulnerability Scanner
Copyright (C) 2016  Future Architect, Inc. Japan.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package gost

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/future-architect/vuls/config"
	"github.com/future-architect/vuls/models"
	"github.com/knqyf263/gost/db"
	"github.com/parnurzeal/gorequest"
)

// Client is the interface of OVAL client.
type Client interface {
	FillWithGost(db.DB, *models.ScanResult) error

	//TODO implement
	// CheckHTTPHealth() error
	// CheckIfGostFetched checks if Gost entries are fetched
	// CheckIfGostFetched(db.DB, string, string) (bool, error)
	// CheckIfGostFresh(db.DB, string, string) (bool, error)
}

// NewClient make Client by family
func NewClient(family string) Client {
	switch family {
	case config.RedHat, config.CentOS:
		return RedHat{}
	case config.Debian:
		return Debian{}
	default:
		return Pseudo{}
	}
}

// Base is a base struct
type Base struct {
	family string
}

// CheckHTTPHealth do health check
func (b Base) CheckHTTPHealth() error {
	if !b.isFetchViaHTTP() {
		return nil
	}

	url := fmt.Sprintf("%s/health", config.Conf.GostDBURL)
	var errs []error
	var resp *http.Response
	resp, _, errs = gorequest.New().Get(url).End()
	//  resp, _, errs = gorequest.New().SetDebug(config.Conf.Debug).Get(url).End()
	//  resp, _, errs = gorequest.New().Proxy(api.httpProxy).Get(url).End()
	if 0 < len(errs) || resp == nil || resp.StatusCode != 200 {
		return fmt.Errorf("Failed to request to OVAL server. url: %s, errs: %v",
			url, errs)
	}
	return nil
}

// CheckIfGostFetched checks if oval entries are in DB by family, release.
func (b Base) CheckIfGostFetched(driver db.DB, osFamily string) (fetched bool, err error) {
	//TODO
	return true, nil
}

// CheckIfGostFresh checks if oval entries are fresh enough
func (b Base) CheckIfGostFresh(driver db.DB, osFamily string) (ok bool, err error) {
	//TODO
	return true, nil
}

func (b Base) isFetchViaHTTP() bool { // Default value of OvalDBType is sqlite3
	return config.Conf.GostDBURL != "" && config.Conf.GostDBType == "sqlite3"
}

// Pseudo is Gost client except for RedHat family and Debian
type Pseudo struct {
	Base
}

// FillWithGost fills cve information that has in Gost
func (pse Pseudo) FillWithGost(driver db.DB, r *models.ScanResult) error {
	return nil
}

func major(osVer string) (majorVersion string) {
	return strings.Split(osVer, ".")[0]
}