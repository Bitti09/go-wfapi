# go-wfapi
WIP  Warframe API Parser with  MQTT publisher

## Current Status:

### Parser:
 - [x] News
 - [x] Sorties
 - [x] Void Fissures
 - [ ] Alerts (waiting for api response)
 - [x] Darvo's Deals ( untranslated )
 - [x] Nightwave
 - [x] Syndicate Missions ( untranslated )
 - [x] Invasions  ( half translated )
 - [ ] Void Trader (waiting for api response)
 - [ ] World Events


### General

- [ ] **Code Rewrite  ( in progress )**
- [ ] **Code Spliting ( in progress )**
- [ ] Rewrite of the  current "hacky" ways for json parsing


## Demo

- Web: https://api.mybitti.de
  - Ping:   /
  - Worldstate(unparsed) /:platform
  - Darvos Deal: /:platform/platform/  (Accept-Language required)
  - News /:platform/news/ (Accept-Language required)
  only first two chars  of  (Accept-Language) is used.


- MQTT Client:  (sometimes offline when i'm testing new updates)
  - Host: mybitti(.)de
  - Port: 1884
  - Protocol: wss
  - Data path: /wf/{lang}/{platform}/
  - Lang: {"en", "de", "es", "fr","it","ko","pl","pt","ru","zh"}
  - Platform: {"pc", "ps4", "xb1", "swi"}
  - Tested Client: [MQTT Explorer](https://mqtt-explorer.com/)

