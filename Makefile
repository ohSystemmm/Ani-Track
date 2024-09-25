#  __  __       _         __ _ _      
# |  \/  | __ _| | _____ / _(_) | ___ 
# | |\/| |/ _` | |/ / _ \ |_| | |/ _ \
# | |  | | (_| |   <  __/  _| | |  __/
# |_|  |_|\__,_|_|\_\___|_| |_|_|\___|
# 
# by ohSystemmm <3

BinaryFile = ani-track
InstallPath = /usr/local/bin/

install:
	go build -o $(BinaryFile)
	sudo install -Dm755 $(BinaryFile) $(InstallPath)$(BinaryFile)

uninstall:
	sudo rm -f $(InstallPath)$(BinaryFile)
