.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	goconvey -cover -port=9090 -workDir="$(realpath .)" -depth=0


.PHONY: goxc
goxc:
	goxc
