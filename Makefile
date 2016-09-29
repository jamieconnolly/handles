NAME = handles

EXECUTABLE = bin/$(NAME)

all: build

build: clean $(EXECUTABLE)

clean:
	@rm -f $(EXECUTABLE)

$(EXECUTABLE):
	go build -o $@ main.go

.PHONY: all build clean
