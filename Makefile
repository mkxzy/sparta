# #####################
# sparta makefile
# written by mkxzy & MoonWolf
# 2018-4-8
# #####################

.PHONY: clean sync

outdir=bin
name=sparta
suffix=.exe

linux: 
	go build -o $(outdir)/$(name)
    
windows: 
	go build -o $(outdir)/$(name)$(suffix)

sync:
	govendor sync

clean:
	rm -rf $(outdir)
