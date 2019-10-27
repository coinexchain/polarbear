package keybase

const (
	errPrefix = "POLARBEARError:"

	errCreate        = "Create key:"
	errDelete        = "Delete key:"
	errRecovery      = "Recovery key:"
	errAdd           = "Add key:"
	errExport        = "Export key:"
	errList          = "List keys:"
	errGetAddress    = "Get address:"
	errGetPubkey     = "Get pubkey:"
	errResetPassword = "Reset password:"
	errGetSigner     = "Get signer:"
	errSign          = "Sign:"
	errSignStdTx     = "Sign stdTx:"
	errSignAndBuild  = "Sign and build bytes for broadcast:"
)

func createKeyErr(err error) string {
	return errPrefix + errCreate + err.Error()
}

func deleteKeyErr(err error) string {
	return errPrefix + errDelete + err.Error()
}

func recoveryKeyErr(err error) string {
	return errPrefix + errRecovery + err.Error()
}

func addKeyErr(err error) string {
	return errPrefix + errAdd + err.Error()
}

func exportKeyErr(err error) string {
	return errPrefix + errExport + err.Error()
}

func listKeysErr(err error) string {
	return errPrefix + errList + err.Error()
}

func getAddressErr(err error) string {
	return errPrefix + errGetAddress + err.Error()
}

func getPubKeyErr(err error) string {
	return errPrefix + errGetPubkey + err.Error()
}

func getSignerErr(err error) string {
	return errPrefix + errGetSigner + err.Error()
}

func resetPasswordErr(err error) string {
	return errPrefix + errResetPassword + err.Error()
}

func signErr(err error) string {
	return errPrefix + errSign + err.Error()
}

func signStdTxErr(err error) string {
	return errPrefix + errSignStdTx + err.Error()
}

func signAndBuildErr(err error) string {
	return errPrefix + errSignAndBuild + err.Error()
}
