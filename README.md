# Go fagkveld
Før workshop bør du installere Go og sjekke at du får kompilert og kjørt et testprogram.

### Sette opp prosjekt
Installer Go
```
brew install go
```
_Legg til støtte for Go i editoren din også!_

Legg til prosjektet (det er **sterkt** anbefalt å legge mappen akkurat her, vi kommer til å snakke mer om det under selve fagkvelden)
```
mkdir -p ~/go/src/github.com/iterate
cd ~/go/src/github.com/iterate
git clone git@github.com:rugern/go-workshop.git
cd go-workshop
```

Installer taskrunner (gir bl.a. automatisk reloading og testing ved endringer av koden) med Go
```
go get github.com/oxequa/realize
```

Installer package manager med brew eller go
```
brew install dep
# eller
go get github.com/tools/godep
```

Installer avhengigheter
```
dep ensure
```

### Kjør testprosjekt
Kompiler kode og kjør binærfil
```
make clean && make
bin/go-workshop
# => "Hello world!"
```

Kjør koden med realize
```
make dev
# => "Hello world!"
```
