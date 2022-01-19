module strongpasswordchecker

go 1.16

replace github.com/strongpasswordchecker/domain => ../strongpasswordchecker/domain

replace github.com/strongpasswordchecker/app => ../strongpasswordchecker/app

require (
	github.com/strongpasswordchecker/app v0.0.0-00010101000000-000000000000
	github.com/strongpasswordchecker/domain v0.0.0-00010101000000-000000000000 // indirect
)
