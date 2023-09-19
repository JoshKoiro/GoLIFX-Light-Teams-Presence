# GoLIFX-Light-Teams-Presence
Golang Implementation of LIFX-Light-Teams-Presence Application - Controlling LIFX light bulb using MS Teams status indicators

# Installation

Download the latest release zip file for your system, unzip the archive. 

enter your Graph API in the `.env` file

Enter your LIFX API key in the `config.yaml` file.

Run the application by double clicking on the icon, or by opening a terminal window and running `./Teams-LIFX`

# Usage

The program acts as a server that reaches out to the graph API and updates your LIFX light every 10 seconds (by default).

Customization of color, brightness and other options may be configured with the `config.yaml` file. The server is reading this file on every tick (10 seconds by default) so changes can be made without having to restart the server application.
