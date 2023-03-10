# VPN_SELLER_INFORMATION

***Author***: Ismail DINC, Quentin VUERICH

***Teachers***: Mr. Alexandre DULAUNOY [@adulau](https://github.com/adulau) & Mr. Christian STUDER [@chris3rd](https://github.com/chrisr3d)

## Context
This repository is related to a university project which is about Threat Intelligence, and Malware Information Sharing Platform (MISP).



## Purpose

The purpose of this project is to integrate two galaxies in MISP:
- VPN Seller List
- IP List for each Seller

## Methodology

The methodology that we are using for this project is the following:

##### Finding sources of information

The first step of our project was to find a lot of different sources about VPN Sellers in order to have consistant and relevant data in our galaxies.
Therefore, we decided to cross-check the web sites to be sure of the information present in our galaxies.
We decided to create an excell file to store the different data such as the name of the VPN, the number of servers, number of locations, number of countries and the different sources (gathering the known IPs provided by these sellers) in ordrer to scrap them later.



##### Scraping data

The second step was to scrap the data from the differents sources we found to format it into json files and to create our galaxies. 
At this moment we decided to create two seperated files that's to say one only for the basic information (name, number of servers, locations and countries) and one for the IPs provided by the sellers in ordrer for the galaxies to be more relevant and exploitable.

To accomplish this task we have created a GO script that will on one hand, parse the data from the excell file previously mentionned and on the other hand scrap the IPs provided by the sellers on the different sources we found.


## Conclusion

This university project was the perfect opportunity to discover a tool that we didn't know before and to contribute by creating galaxies. It allowed us to develop our ability to search relevant information and to learn on the MISP. Finally, we wanted to thanks our teachers for the opportunity to work on this project.
