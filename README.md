# alkalarm WIKI

Read the documentation with the wiki 

https://github.com/alknopfler/alkalarm/wiki

# Welcome to the alkalarm wiki!

**AlkAlarm** is a open source project created to get ready a home Alarm System in few minutes. In the market you will buy a home system alarm based on 433Mhz sensors and using GSM modules to send the notifications like: 
![Alarm System ](https://www.alarmasparacasabaratas.com/wp-content/uploads/2017/08/alarma-gsm-eray-comprar-online.jpg)

The main idea of this project is based on the same kind of Alarm GSM systems, but in this case, we only have to buy the sensors and sirens, because:
* the main controller, and gsm module will be made with the **raspberry** (physically)
* the **AlkAlarm** code of this repo will be installed on raspberry and it will be the brain of raspberry Alarm System
* Sensor bought will be **generic 433Mhz** and will be added to the raspberry using a 433mhz receptor.
* **Notifications** will be send by email to get it free, and wired/wireless sirens will be connected to Raspberry.
* **WebServer**, **API server**, and **Android APP** will be available to communicate with the Alarm System.

# Main Features
* Interface with the generic 433Mhz security sensors in order to register into the system.
* Detect sensor events registered previously and trigger the notifications actions.
* Operate with remote controllers at the same level than the webserver, api server and android app.
* Webserver, API and Android APP integration for management the system.
* Notifications by email (without any limit) and triggering sirens connected to the system.
* Basic Auth in all the operations using Access Code
* Discovering the 433Mhz sensors to add new sensors in the future
* Zones in order to separate the sensor in rooms


Read the documentation with the wiki 

https://github.com/alknopfler/alkalarm/wiki