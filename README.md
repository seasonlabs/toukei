# Toukei is a highly-concurrent, multithreaded, realtime, git repository statistics server written in Go (golang)

Use toukei to offer an insight of your work or have a bird's eye over your company git stats. For example we at http://www.season.es use toukei to provide website visitors with internal git commits and files count.

## Requirements:

* Git
* Redis
* Go >= 1.0.0

## Installation

Clone the repository

Build with: 
	
	./all.sh

## Deploy 

Copy bin directory to a server with git repositories. Use a upstart/init.d script to start the service

## Configure

Write the path to a git repositories dir in bind/config.yml

## Demo

http://toukei.season.es:12345/

