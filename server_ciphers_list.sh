#!/usr/bin/env bash

nmap -sV --script ssl-enum-ciphers -p 8443 localhost
