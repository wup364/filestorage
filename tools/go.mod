module fstools

go 1.14

// replace github.com/wup364/pakku => ../../pakkuboot/pakku
replace github.com/wup364/filestorage/opensdk => ../opensdk

require (
	github.com/go-git/go-billy/v5 v5.3.1
	github.com/mattn/go-colorable v0.1.12
	github.com/willscott/go-nfs v0.0.0-20211118152618-00ba06574ea0
	github.com/wup364/filestorage/opensdk v0.0.1
	github.com/wup364/pakku v0.0.1
	golang.org/x/net v0.0.0-20220930213112-107f3e3c3b0b
)
