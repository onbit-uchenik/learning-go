module example.com/hello

go 1.16

replace example.com/fibonacci => ../fibonacci

require example.com/fibonacci v0.0.0-00010101000000-000000000000 // indirect
