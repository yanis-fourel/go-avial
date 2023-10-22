enumstring: ./hgtp/taxonomy.go taxonomy.go
	stringer -type=Command,Error,Report -linecomment hgtp/taxonomy.go && mv hgtp/command_string.go hgtp/taxonomy_string.go
	stringer -type=Class,Action,Aspect,Category,Parameter,Attribute,Tag,Goal,Mode,Unit,Event,Match,State,Format,Method,Notice,Prefix,Scheme,Context,Percept,Constant,Language,Operator,Platform,Template,Condition,Permission,Precedence,Installation -linecomment taxonomy.go && mv class_string.go taxonomy_string.go

test:
	go test github.com/yanis-fourel/go-avial/Go \
		github.com/yanis-fourel/go-avial/Go/hgtp_tests
