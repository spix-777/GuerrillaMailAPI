# Guerrilla Mail JSON API

This is a Go program that interacts with the Guerrilla Mail JSON API. It allows you to generate a random email address and check for new emails.

## Overview

The Guerrilla Mail JSON API provides a simple way to generate temporary email addresses and retrieve emails sent to those addresses. This program demonstrates how to use the API to perform these actions.

## Features

- Generate a random email address.
- Check for new emails in the generated mailbox.
- Print the details of received emails.

## Usage

1. Clone the repository:

   ```shell
   git clone https://github.com/spix-777/GuerrillaMailAPI

2. Build the program:

   ``` shell
   Copy code
   go build

3. Run the program:

   ``` shell
   Copy code
   ./guerrilla-mail-json-api

The program will display a banner and generate a random email address using the Guerrilla Mail JSON API. It will then continuously check for new emails and print their details when received.
Press Ctrl+C to stop the program and exit.

Command-line Options

The program supports the following command-line options:

-v: Prints the version information and exits.
Requirements

Go 1.15 or later.
API Reference

The program utilizes the Guerrilla Mail JSON API. For more information about the API, refer to the official documentation.

License

This code is licensed under the MIT License.

Credits

Made by: SpiX-777
Date: 2023-06-13
Version: 1.0.0