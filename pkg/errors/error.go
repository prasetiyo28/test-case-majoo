package errors

import "errors"

// New returns an error that formats as the given text.
var New = errors.New

// known errors
var (
	ErrLogin               = errors.New("Username atau password tidak sesuai.")
	ErrClaimAccessToken    = errors.New("Gagal generate Access token.")
	ErrClaimRefreshToken   = errors.New("Gagal generate refresh token.")
	ErrClaimToken          = errors.New("Gagal generate token.")
	ErrPublicKey           = errors.New("Tidak Mengenali Public Key.")
	ErrFormatRequestBody   = errors.New("Format data yang dikirim tidak valid.")
	ErrInvalidToken        = errors.New("Token tidak valid")
	ErrParamInvalid        = errors.New("Parameter tidak valid.")
	ErrInternResp          = errors.New("Terdapat kesalahan pada internal")
	ErrAuthorizationBearer = errors.New("Harus memberikan header otorisasi dengan format `Bearer {token}`")
)
