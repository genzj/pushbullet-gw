ifdef SYSTEMROOT

ifdef SHELL
  # in MINGW
  RM = rm -f
else
  RM = del /Q
endif

  BINARY_SUFFIX = .exe
  FixPath = $(subst /,\,$1)
else
   ifeq ($(shell uname), Linux)
      RM = rm -f
      FixPath = $1
   endif
  BINARY_SUFFIX =
endif

BINARY_NAME = pushbullet-gw$(BINARY_SUFFIX)

.PHONY: all
all: $(BINARY_NAME)
	rice --import-path "github.com/genzj/pushbullet-gw/server" append --exec "$(BINARY_NAME)"

$(BINARY_NAME):
	go build -ldflags "-s" -o "$(BINARY_NAME)"

.PHONY: clean
clean:
	$(RM) $(BINARY_NAME)

