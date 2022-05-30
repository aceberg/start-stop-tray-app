# start-stop-tray-app
Tray icon for my cli app


Create go modules from .png icons:
```
bash make_icon.sh location.png
bash make_icon.sh location1.png
```

Build:
```
go mod init start-stop-tray-app
go mod tidy
go build
```

Example and make_icon.sh from:
https://github.com/getlantern/systray

Icons (location):
https://www.flaticon.com/free-icon/location_535137
