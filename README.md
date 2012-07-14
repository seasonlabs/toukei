# Toukei is a highly-concurrent, multithreaded, realtime, git repository statistics server written in Go (golang)

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

