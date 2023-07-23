module fstools

go 1.14

// replace github.com/wup364/pakku => ../../pakkuboot/pakku
replace github.com/wup364/filestorage/opensdk => ../opensdk

require (
	github.com/mattn/go-colorable v0.1.12
	github.com/wup364/filestorage/opensdk v0.0.1
	github.com/wup364/pakku v0.0.5
	golang.org/x/sys v0.0.0-20220728004956-3c1f35247d10 // indirect
)
