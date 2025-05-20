# echocat's Home Assistant Add-ons Repository

## About

Home Assistant allows anyone to create add-on repositories to share their add-ons for Home Assistant easily. This repository is one of those repositories, providing extra Home Assistant add-ons for your installation.

The primary goal of this project is to provide you (as a Home Assistant user) with additional, high quality, add-ons that allow you to take your automated home to the next level.

## Installation

[![Add repository on my Home Assistant][repository-badge]][repository-url]

[repository-badge]: https://img.shields.io/badge/Add%20repository%20to%20my-Home%20Assistant-41BDF5?logo=home-assistant&style=for-the-badge
[repository-url]: https://my.home-assistant.io/redirect/supervisor_add_addon_repository/?repository_url=https%3A%2F%2Fgithub.com%2Fechocat%2Fhassos-addons

If you want to do add the repository manually, please follow the procedure highlighted in the [Home Assistant website](https://home-assistant.io/hassio/installing_third_party_addons). Use the following URL to add this repository: https://github.com/echocat/hassos-addons

## Add-ons provided by this repository

### [Duplicati](https://github.com/echocat/hassos-addon-duplicati)
![Version](https://img.shields.io/github/release/echocat/hassos-addon-duplicati.svg?label=Version&style=flat-square)
![Ingress](https://img.shields.io/badge/dynamic/yaml?label=Ingress&style=flat-square&query=%24.ingress&url=https%3A%2F%2Fraw.githubusercontent.com%2Fechocat%2Fhassos-addon-duplicati%2Fmain%2Fconfig%2Fconfig.yaml)
![Arch](https://img.shields.io/badge/dynamic/yaml?color=success&label=Arch&style=flat-square&query=%24.arch&url=https%3A%2F%2Fraw.githubusercontent.com%2Fechocat%2Fhassos-addon-duplicati%2Fmain%2Fconfig%2Fconfig.yaml)

If you want to create more precise backups of your Home Assistant than what the built-in tools offer, [Duplicati](https://duplicati.com/) is a helpful solution. It allows you to back up either the entire system or selected files locally or to a wide range of cloud storage providers.

