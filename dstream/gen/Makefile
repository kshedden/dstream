all: bcols base comparisons segment utils convert regroup dechunk replace generate memcopy streamcsv filtercol dataframe

.PHONY: bcols base comparisons segment utils convert regroup dechunk replace generate memcopy streamcsv filtercol dataframe

bcols:
	go run main.go -template=bcols.tmpl
segment:
	go run main.go -template=segment.tmpl
utils:
	go run main.go -template=utils.tmpl
comparisons:
	go run main.go -template=comparisons.tmpl
convert:
	go run main.go -template=convert.tmpl -numeric
regroup:
	go run main.go -template=regroup.tmpl
dechunk:
	go run main.go -template=dechunk.tmpl
replace:
	go run main.go -template=replace.tmpl
generate:
	go run main.go -template=generate.tmpl
memcopy:
	go run main.go -template=memcopy.tmpl
streamcsv:
	go run main.go -template=streamcsv.tmpl
filtercol:
	go run main.go -template=filtercol.tmpl
dataframe:
	go run main.go -template=dataframe.tmpl
