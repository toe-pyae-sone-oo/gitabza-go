package artistutil

var SupportedArtistPicFiles = []string{"jpg", "jpeg", "png", "svg"}

func IsFileSupported(fileExt string) bool {
	for _, ext := range SupportedArtistPicFiles {
		if ext == fileExt {
			return true
		}
	}
	return false
}
