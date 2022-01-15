package utils

import (
	"cloud.google.com/go/firestore"
)

const (
	FIRESTORE_TIME = firestore.ServerTimestamp

	USER_COLLECTION                = "user_kompas"
	EVENT_COLLECTION               = "eventKompas"
	EVENT_COLLECTION_ATTENDEEE     = "eventAttendee"
	USER_COLLECTION_WHITELIST      = "whitelist"
	COUNTRY_COLLECTION             = "country"
	ADDRESS_COLLECTION             = "addresses"
	GOOGLE_SUBSCRIPTION_COLLECTION = "google-subs"
	NOTIFICATION_COLLECTION        = "shannonLogs"
	USER_ACCESS_COLLECTION         = "userAccess"
	QRC_NEWSPAPER_COLLECTION       = "qrcNewsPaper"
	QRC_LOG_COLLECTION             = "qrcNewsPaperLogs"

	FIRESTORE_PROJECT_ID       = "kmn-dev"
	FIRESTORE_PROJECT_ID_EVENT = "event-kompasid-dev"

	PASSWORD_WEAK   = "lemah"
	PASSWORD_MEDIUM = "sedang"
	PASSWORD_STRONG = "kuat"

	TEMPLATE_MAIL_VERIFICATION       = "verifikasi"
	TEMPLATE_MAIL_RESET_PASSWORD     = "resetPassword"
	TEMPLATE_MAIL_RESET_PASSWORD_VIP = "resetPasswordVIP"
	TEMPLATE_MAIL_LENGKAPI_DATA      = "lengkapiDataDiri"
	TEMPLATE_MAIL_REGISTRASI_VIP     = "registerVIP"

	SUBJECT_MAIL_VERIFICATION   = "Verifikasi Akun Kompas.id"
	SUBJECT_MAIL_RESET_PASSWORD = "Atur Ulang Kata Sandi Akun Kompas.id"
	SUBJECT_MAIL_LENGKAPI_DATA  = "lengkapi Data Diri"
	SUBJECT_REGISTRASI_VIP      = "Akun Kompas.id Berhasil Didaftarkan"

	TOKEN_TYPE_ACCESS  = 1
	TOKEN_TYPE_REFRESH = 2
	YEAR_INT           = 31556952

	SSO_CLIENT_COLLECTION = "ssoClient"

	TYPE_MAIL_VERIFICATION   = "verify email"
	TYPE_MAIL_RESET_PASSWORD = "reset password"
	TYPE_MAIL_REGISTRASI_VIP = "registrasi vip"
	TYPE_USER_NON_VIP        = "non vip"
	TYPE_USER_VIP            = "vip"
	TYPE_MAIL_BROADCAST      = "broadcast email"

	TYPE_MAIL_LENGKAPI_DATA = "lengkapi data diri"

	GENDER_MALE   = 1
	GENDER_FEMALE = 2

	JSON_CONTENT_TYPE = "application/json"

	VERIFICATION_NOTIF_DESCRIPTION = "Jika belum menerima tautan verifikasi, klik di sini untuk ajukan kirim ulang tautan verifikasi ke email Anda."
	RESET_PASS_NOTIF_DESCRIPTION   = "Atur ulang kata sandi telah dikirimkan ke email"

	VIP_DEFAULT_PASSWORD      = "kompas1234"
	LENGKAPI_DATA_DESCRIPTION = "Lengkapi data diri Anda untuk mendapatkan pengalaman terbaik di Kompas.id."

	QONTAK_TEMPLATE_ID         = "285b19b4-22a1-420e-9bcb-ba9104cd6064"
	QONTAK_CHANNEL_INTEGRATION = "22a75949-446a-4267-8ac2-1f105441d8fc"
	QONTAK_DIRECT_PATH         = "/api/open/v1/broadcasts/whatsapp/direct"
	QONTAK_OAUTH_PATH          = "/oauth/token"
	QONTAK_CREATE_LIST         = "/api/open/v1/contacts/contact_lists"
	QONTAK_SEND_BROADCAST      = "/api/open/v1/broadcasts/whatsapp"

	// Membership rabbitMQ
	SOURCE   = "other"
	IP       = "127.0.0.1"
	PRIORITY = 1

	BROADCAST_TYPE_WHATSAPP     = 1
	BROADCAST_TYPE_EMAIL        = 2
	BROADCAST_TYPE_NOTIFICATION = 3

	SUBJECT_MEMBERSHIP_SHORTENED = "Akses Membership Diperpendek"
	SUBJECT_MEMBERSHIP_REVOKED   = "Akses Membership Dihentikan"

	TEMPLATE_MEMBERSHIP_SHORTENED = "shortened_membership"
	TEMPLATE_MEMBERSHIP_REVOKED   = "revoked_membership"

	TEMPLATE_TYPE_BROADCAST_SHORTENED = 1
	TEMPLATE_TYPE_BROADCAST_REVOKED   = 2

	OTP_LENGTH    = 6
	FLAG_REGISTER = 1
	FLAG_LOGIN    = 2

	TYPE_EMAIL        = "email"
	TYPE_PHONE_NUMBER = "phoneNumber"
)
