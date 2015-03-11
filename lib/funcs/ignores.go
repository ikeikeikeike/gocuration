package funcs

func IsEmbed(domain string) bool {
	switch domain {
	case "japan-whores.com":
		return false
	case "pornhub.com":
		return false
	case "xhamster.com":
		return false
	case "redtube.com":
		return false
	case "tube8.com":
		return false
	default:
		return true
	}
}
