package middlewares

import (
	"strings"

	"github.com/GooDu-dev/gd-practical-project-backend/src/v1/common"
	settings "github.com/GooDu-dev/gd-practical-project-backend/utils"
	devfleError "github.com/GooDu-dev/gd-practical-project-backend/utils/error"
)

type ValidatorService struct {
	BasicHeader HeaderRequest
	UserHeader  UserHeaderRequest
}

func (h *HeaderRequest) CheckContentType() (err error) {
	if common.IsDefaultValueOrNil(h.ContentType) {
		return devfleError.MissingRequestError
	}
	if content_type, err := settings.ContentType.Value(); err == nil {
		if h.ContentType == content_type {
			return nil
		}
		return devfleError.InvalidHeaderNotAcceptableError
	}
	return err
}

func (h *HeaderRequest) CheckContentCode() (err error) {
	// public key check
	if common.IsDefaultValueOrNil(h.ContentCode) {
		return devfleError.MissingRequestError
	}
	if content_code, err := settings.ContentCode.Value(); err == nil {
		if h.ContentCode == content_code {
			return nil
		}
		return devfleError.InvalidHeaderNotAcceptableError
	}
	return nil
}

func (h *HeaderRequest) CheckClientVersion() (version string, err error) {
	// check web version
	lst := strings.Split(h.ClientVersion, ".")
	version = lst[0]
	if v, err := settings.ClientVersion.Value(); err != nil {
		if version == v {
			return version, nil
		}
	}
	return "", devfleError.InvalidHeaderNotAcceptableError
}

func (h *HeaderRequest) CheckAccessCtrl() (err error) {
	// check user token

	return nil
}

func (h *HeaderRequest) CheckSourceCtrl() (err error) {
	// check api toekn
	return nil
}
