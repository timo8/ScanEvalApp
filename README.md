# ScanEvalApp

## Úvod

Zadanie na predmete Tímový projekt, ktorého cieľom bolo vytvoriť aplikáciu na zvýšenie rýchlosti a efektivity vyhodnocovania prijímacích skúšok.

Aplikácia umožňuje automatizovať celý proces od vytvorenia a generovania odpoveďových hárkov až po spracovanie naskenovaných testov, vyhodnotenie odpovedí a zobrazenie výsledkov. Pri spracovaní sa využíva OCR, počítačové videnie a detekcia QR kódov.

## Funkcionalita

- správa študentov a prijímacích skúšok,
- generovanie personalizovaných odpoveďových hárkov,
- spracovanie naskenovaných PDF dokumentov,
- automatická detekcia a vyhodnotenie odpovedí,
- rozpoznávanie identifikačných údajov pomocou OCR,
- detekcia QR kódov pre identifikáciu hárkov,
- generovanie výsledkov a štatistík,
- vytváranie PDF dokumentov a reportov.

## Použité technológie

- **Go** – hlavný programovací jazyk aplikácie,
- **GoCV / OpenCV** – spracovanie obrazu, detekcia odpovedí a vizuálnych prvkov,
- **Tesseract OCR** – rozpoznávanie textu zo skenovaných dokumentov,
- **GORM** – ORM pre prácu s databázou,
- **SQLite** – databáza aplikácie,
- **LaTeX** – generovanie odpoveďových hárkov a PDF dokumentov,
- **Gio** – grafické používateľské rozhranie desktopovej aplikácie.

## Spracovanie testov

Hlavnou časťou projektu je automatizované spracovanie naskenovaných odpoveďových hárkov. Aplikácia najskôr spracuje jednotlivé stránky dokumentu a pomocou počítačového videnia vykoná ich zarovnanie a detekciu potrebných oblastí. Následne sa pomocou OCR získajú identifikačné údaje a pomocou analýzy obrazu sa vyhodnotia označené odpovede.

Jednotlivé stránky je možné spracovávať paralelne, čo umožňuje skrátiť čas potrebný na vyhodnotenie väčšieho množstva testov.

## Cieľ projektu

Cieľom projektu bolo vytvoriť riešenie, ktoré zjednoduší prácu pri vyhodnocovaní prijímacích skúšok, zníži množstvo manuálnej práce a umožní rýchlejšie a spoľahlivejšie získanie výsledkov.



# ScanEvalApp

Dokumentácia dostupná [TU](https://docs.google.com/document/d/1oPEVyG-Ius-a9JKvhcH9mh4ZzbzJkZ4PRGxit0UCV0w/edit?usp=sharing)

### Install dependencies
- nainstaluje dependecy na kompilaciu latex sablony (treba spravit `chmod +x installDependencies.sh`)
- nainstaluje aj golang verziu 1.23.2 (ked to budete spustat druhy a viac krat tak dajte -f aby vam znova nepridavalo cestu do .bashrc)

### Pullnutie novej verzie
```sh
git pull
``` 

### Spustenie (v pripade nutnosti buildovania novej verzie ju zbuilduje a spusti)
```sh
./start.sh
``` 
