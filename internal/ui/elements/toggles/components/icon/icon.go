package icon

import uiutils "github.com/fastygo/templ/utils"

func iconClasses(p IconProps) string {
	sizeClass := "h-4 w-4"
	switch p.Size {
	case "xs":
		sizeClass = "h-3 w-3"
	case "sm", "":
		sizeClass = "h-4 w-4"
	case "md":
		sizeClass = "h-5 w-5"
	case "lg":
		sizeClass = "h-6 w-6"
	}
	return uiutils.Cn("inline-block shrink-0", "latty", "latty-"+p.Name, sizeClass, p.Class)
}
