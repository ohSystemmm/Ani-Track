#  __  __       _         __ _ _      
# |  \/  | __ _| | _____ / _(_) | ___ 
# | |\/| |/ _` | |/ / _ \ |_| | |/ _ \
# | |  | | (_| |   <  __/  _| | |  __/
# |_|  |_|\__,_|_|\_\___|_| |_|_|\___|
# 
# by ohSystemmm <3 

Binary = ani-track
InstallPath = /usr/bin/

install:
	go build -o ani-track
	sudo mv ani-track $(InstallPath)
	

uninstall: 
	sudo rm -rf /usr/bin/ani-track
