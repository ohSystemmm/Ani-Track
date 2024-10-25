#  __  __       _         __ _ _      
# |  \/  | __ _| | _____ / _(_) | ___ 
# | |\/| |/ _` | |/ / _ \ |_| | |/ _ \
# | |  | | (_| |   <  __/  _| | |  __/
# |_|  |_|\__,_|_|\_\___|_| |_|_|\___|
# 
# by ohSystemmm <3

BinaryFile = ani-track
InstallPath = /usr/local/bin/

all: $(BinaryFile)

$(BinaryFile):
	go build -o $(BinaryFile)

install: $(BinaryFile)
	install -Dm755 $(BinaryFile) $(InstallPath)$(BinaryFile)

uninstall:
	@echo "Uninstalling $(BinaryFile) from $(InstallPath)"
	sudo rm -f $(InstallPath)$(BinaryFile)

clean:
	rm -f $(BinaryFile)

.PHONY: all install uninstall clean