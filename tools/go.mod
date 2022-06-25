module fstools

go 1.14

// replace github.com/wup364/pakku => ../../pakkuboot/pakku
replace opensdk => ../opensdk

require (
	github.com/mattn/go-colorable v0.1.12
	github.com/wup364/pakku v0.0.1
	opensdk v0.0.0-00010101000000-000000000000
)
