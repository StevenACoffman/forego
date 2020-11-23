# forego
Simple animated gif desktop display using fyne

# Install and run it

```
# Build somewhere in your PATH that is writeable
go install github.com/StevenACoffman/forego
forego https://github.com/StevenACoffman/forego/blob/main/circle.gif?raw=true
```
Also works with local files:
```
forego turtle-small.gif
```

# Credit where credit is due

The code is mostly from [Andrew Williams](https://github.com/andydotxyz) and [Oliver Eilhard](https://github.com/olivere) (specifically [imgcat](https://github.com/olivere/iterm2-imagetools/tree/master/cmd/imgcat))

I got the `circle.gif` from the [plymouth themes](https://github.com/adi1090x/plymouth-themes) which are a port of some android bootanimations from a forum post [here](https://forum.xda-developers.com/android/themes/alienware-t3721978), so I don't know who the original artist was.

I got the awkward turtle too many years ago to remember where.

Happy to add more specific credit if anyone knows who originated either! 