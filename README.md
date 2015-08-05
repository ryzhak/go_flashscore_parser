# go_flashscore_parser
HTML parser of the http://www.flashscore.com/ using golang

### How to use
* download chromeDriver, for example from here [https://sites.google.com/a/chromium.org/chromedriver/downloads](https://sites.google.com/a/chromium.org/chromedriver/downloads)

### Example
```
package main

import (
	flashscoreParser "github.com/ryzhak/go_flashscore_parser"
)

func main() {
    //path to chromedriver
	chromeDriverPath := "/usr/local/bin/chromedriver"
	games := flashscoreParser.GetGames(chromeDriverPath)
	flashscoreParser.Show(games)
}
```
### Output

```
EUROPE Champions League - Qualification
17:00  FC Astana (Kaz) 3:2 HJK (Fin)
19:30 Qarabag (Aze) : Celtic (Sco)
19:45 Sparta Prague (Cze) : CSKA Moscow (Rus)
20:30 BATE (Blr) : Videoton (Hun)
21:15 Basel (Sui) : Lech Poznan (Pol)
21:30 Club Brugge (Bel) : Panathinaikos (Gre)
21:30 Malmo FF (Swe) : Salzburg (Aut)
21:30 Partizan (Srb) : Steaua Bucuresti (Rou)
21:45 Plzen (Cze) : Maccabi Tel Aviv (Isr)
21:45 Shakhtar (Ukr) : Fenerbahce (Tur)
21:45 Skenderbeu (Alb) : Milsami (Mda)
ARGENTINA Copa Argentina
21:00 Ferro : Los Andes
ASIA East Asian Championship
13:20 Japan 1:1 South Korea
16:00 China 2:0 North Korea
AUSTRALIA FFA Cup
13:30 Croydon Kings 1:2 Queensland Lions 
13:30 Darwin Olympic 1:6 Adelaide United 
13:30  Rockdale City Suns 3:1 Perth SC
13:30 Sorrento 0:2 Sydney FC 
BRAZIL Série B
03:00 America MG 2:0 Parana
03:00  Sampaio Correa 2:0 Bragantino
CHILE Chilean Cup
18:00 Rangers 0:3 U. De Chile
20:00 Coquimbo : La Serena
21:30 Everton : S. Wanderers
ENGLAND WSL 1 Women
21:30 Notts County W : Bristol Academy W
ESTONIA Estonian Cup
19:00 Jarva-Jaani : Vutiselts
19:00 Tabivere : Pirita
19:00 Tammeka Tartu : Maardu
20:30 Tallinna SK Dnipro : Welco Elekter
FINLAND Ykkonen
18:30 JJK Jyväskylä 0:0 Jazz Pori
FINLAND Kakkonen North
18:30 JBK 0:0 YPA
18:30 Kiisto 0:0 Kerho 07
18:30 KPV Kokkola 1:0 GBK Kokkola
18:30 OPS 0:0 Santa Claus
18:30 TP-47 0:0 AC Kajaani
FINLAND Kakkonen South
18:30 Gnistan 1:0 JaPS
18:30 Keski-Uusimaa 0:0 Vaajakoski
18:30 NJS 1:0 TPV
FINLAND Kakkonen East
18:30 Kultsu 0:1 Viikingit
18:30 Lahti Akatemia 0:0 JIPPO
18:30 Sudet 0:0 Klubi 04
FINLAND Kakkonen West
18:30 Åbo 0:0 MuSa
18:30 ESC 0:0 MaPS
18:30 KaaPo 0:1 BK-46
18:30 SalPa 1:1 Narpes
GERMANY Regionalliga Nordost
19:00 Jena : Schönberg
HUNGARY Hungarian Cup
18:00 Babocsa : Mateszalkai MTK
18:00 Nagyecsed : REAC
18:00 Tiszakanyar : Nyirbatori FC
ICELAND Pepsideild
21:00 ÍBV Vestmannaeyjar : Fylkir
22:15 Breidablik : Keflavik
22:15 Fjolnir : KR Reykjavik
22:15 Hafnarfjordur : Valur
22:15 Leiknir : Stjarnan
22:15 Vikingur Reykjavik : Akranes
ISRAEL Toto Cup
19:00 Sakhnin : Maccabi Haifa
19:30 Hapoel Kfar-Saba : H. Raanana
20:00 Hapoel Haifa : H. Akko
LATVIA Latvian Cup
19:00 Preilu BJSS : Spartaks
LITHUANIA A Lyga
18:00 Stumbras 0:1 Atlantas
20:00 Siauliai : Trakai
LITHUANIA Lithuanian Cup
18:00 Rotalis : Silute
MALAYSIA Super League
15:45 FELDA 0:0 Kelantan
15:45 Johor DT 4:0 Sime Darby
15:45 Pahang 2:0 ATM FA
15:45 PDRM FA 3:2 Perak
15:45 Sarawak FA 1:1 Selangor
15:45 Terengganu 4:2 LionsXII
MALTA Summer Cup
21:15 Balzan Youths : Qormi
MEXICO Copa Mexico - Apertura
03:00 Cruz Azul 0:1 Venados 
03:00 Monarcas 1:0 Dep. Tepic
03:00 Puebla 2:0 Celaya 
03:00 Zacatepec 1:2 Club Tijuana
05:00 Atlante 3:2 Pachuca
05:00 Murcielagos 1:0 Dorados de Sinaloa
05:00 Veracruz 2:1 Lobos BUAP
NIGERIA Premier League
18:00 Abia Warriors 2:0 Giwa
18:00 Akwa 1:1 Enyimba
18:00 Enugu 0:0 Taraba
18:00 Kano 0:1 Heartland
18:00 Kwara 0:1 Dolphins
18:00 Lobi 2:0 Bayelsa
18:00 Nasarawa 2:1 El Kanemi
18:00 Shooting 0:1 Wikki
18:00 Warri 1:0 Ifeanyi Ubah
NORTH & CENTRAL AMERICA CONCACAF Champions League
03:00 Querétaro 2:0 San Francisco
05:00 Municipal 0:1 Real Salt Lake
05:00 Santos Laguna 4:0 W Connection
NORTH & CENTRAL AMERICA International Champions Cup
22:00 Chelsea : Fiorentina
PHILIPPINES UFL
12:15 Pachanga 1:3 Stallion
14:45  Manila Jeepney 2:1 Ceres
SINGAPORE S.League
15:15 Brunei DPMM 3:1 Geylang
SLOVAKIA Slovak Cup
19:00 Rudina : Oravské Veselé
SOUTH AFRICA MTN 8 Cup
20:30 Kaizer Chiefs : Maritzburg Utd
20:30 Mamelodi Sundowns : Bloem Celtic
SWEDEN Division 2 - Södra Götaland
20:00 Asarums : Nosaby IF
20:00 Prespa Birlik : Lindsdals
SWEDEN Svenska Cupen - Qualification
19:00 Husie : Hollvikens
19:30 Upsala : Akropolis
20:00 Melleruds : Carlstad
20:00 Nybro : Oskarshamns
20:00 Ostersund IFK : Harnosands
20:00 Sandvikens : Brage
20:00 Sollentuna : Vasalunds
20:00 Vara SK : Norrby
20:30 Nacka FF : Huddinge
20:45 Gute : Sodertalje FK
SWEDEN Allsvenskan Women
20:00 Eskilstuna United W : Rosengard W
USA USL
03:00 Oklahoma City Energy 2:0 Los Angeles 2
VENEZUELA Copa Venezuela
22:30 Atl. Socopo : Zamora
VIETNAM V-League
14:00 Than Quang Ninh 3:0 Gia Lai 
WORLD Club Friendly
18:00 Zawisza (Pol) 0:1 Anorthosis (Cyp)
18:30 Torino (Ita) 0:0 Pro Vercelli (Ita)
19:00 Al-Sadd (Qat) : Al Ain (Uae)
19:00 Entella (Ita) : Genoa (Ita)
19:00 Legnica (Pol) : AEK (Gre)
19:00 Veria (Gre) : Skoda Xanthi (Gre)
19:00 Verona (Ita) : Al-Hilal (Sau)
19:30 Tudelano (Esp) : Ebro (Esp)
20:00 Iraklis (Gre) : Panthrakikos (Gre)
20:00 Murcia (Esp) : Elche (Esp)
20:00 Tondela (Por) : Berkane (Mar)
21:00 Estoril (Por) : Setubal (Por)
21:00 Rio Ave (Por) : Valladolid (Esp)
21:30 Guijuelo (Esp) : UD Logrones (Esp)
21:30 Leganes (Esp) : Rayo Vallecano (Esp)
21:30 Ponferradina (Esp) : Gijon (Esp)
21:30 R. Oviedo (Esp) : Dep. La Coruna (Esp)
21:30 Udinese (Ita) : Spal (Ita)
21:45 Buxton (Eng) : Sheffield Utd U21 (Eng)
21:45 Pisa (Ita) : Empoli (Ita)
23:00 Barcelona (Esp) : AS Roma (Ita)
WORLD Audi Cup
19:15 Tottenham : AC Milan
21:45 Bayern Munich : Real Madrid
WORLD Friendly International Women
19:00 Slovakia W : United Arab Emirates W
```
