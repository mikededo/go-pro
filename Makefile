DOCS=list/linked_list/README.md list/stack/README.md tree/bitree/README.md

all: $(DOCS)

%/README.md: %
	gomarkdoc --output $@ ./$<

