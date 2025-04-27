# Installation

go mod tidy<br>
go build -o go_player && ./go_player "example.mp3"<br>
For Linux : You need, pipewire pipewire-alsa pipewire-pulse packages.
Redirect Alsa configuration to PipeWire.<br>

### EXAMPLE :
```bash
mkdir -p ~/.config/alsa/conf.d
echo -e 'pcm.!default {\n  type pipewire\n}\nctl.!default {\n  type pipewire\n}' | sudo tee /etc/alsa/conf.d/99-pipewire-default.conf
```

# Usage (Linux)

./go_player "path-to-mp3"

### EXAMPLE :
```bash
./go_player "/media/mal1kore1s/Data1/Sounds & Musics/Music/Nope your too late i already died - wifiskeleton.mp3"
```

## Todo

**Volume Changer**<br>
**Finish fix**<br>
**More song support (opening more than 1 file)**<br>
**Equalizer (maybe i'm not sure :d)**<br>