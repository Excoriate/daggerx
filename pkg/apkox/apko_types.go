package apkox

// KeyringSkeleton represents the structure of a keyring entry.
type KeyringSkeleton struct {
	// Path is the file system path to the keyring file.
	// It may be empty if only a URL is provided.
	Path string
	// URL is the web address from which the keyring can be downloaded.
	URL string
}
