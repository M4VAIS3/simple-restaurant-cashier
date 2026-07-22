package utils

import (
	"fmt"
	"html/template"
	"strings"
)

// TemplateFuncs berisi fungsi-fungsi helper untuk template HTML
var TemplateFuncs = template.FuncMap{
	// inc: increment index (untuk nomor urut di tabel)
	"inc": func(i int) int {
		return i + 1
	},
	// formatRupiah: format angka ke format Rupiah dengan titik
	"formatRupiah": func(n int) string {
		s := fmt.Sprintf("%d", n)
		result := ""
		for i, c := range s {
			if i > 0 && (len(s)-i)%3 == 0 {
				result += "."
			}
			result += string(c)
		}
		return result
	},
	// shortDate: format singkat tanggal
	"shortDate": func(s string) string {
		return strings.TrimSpace(s)
	},
}
