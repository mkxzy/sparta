# #####################
# sparta makefile
# written by mkxzy & MoonWolf
# 2018-4-8
# #####################

.PHONY: clean sync gen

outdir=bin
name=sparta

build: 
	go build -o $(outdir)/$(name)

sync:
	govendor sync

clean:
	rm -rf $(outdir)
