package utils

import uiutils "github.com/fastygo/templ/utils"

// Cn joins class fragments (delegates to github.com/fastygo/templ/utils).
func Cn(classes ...string) string {
	return uiutils.Cn(classes...)
}
