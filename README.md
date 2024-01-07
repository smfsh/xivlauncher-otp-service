# XIVLauncher OTP Service

## Overview
The XIVLauncher OTP Service is a Go application designed to generate and send One-Time Passwords (OTPs) to a specified set of IP addresses running the XIVLauncher service. It periodically checks if the service is available on each IP and, upon successful detection, generates an OTP and sends it to the corresponding service endpoint.

This service exists simply as a ridiculous quality-of-life feature to save the approximately four seconds of time it would take to type your six digit OTP. Anyone using it should immediately evaluate their life choices. Also, please understand what this service does. It spams your network with traffic, it potentially exposes your secrets if you're not careful, and it uses a non-zero amount of server resources to exist.

## Features
- **Multiple IP Support**: Can handle one or more IP addresses.
- **Environment Variable & Command-Line Argument Support**: Accepts input either through environment variables or command-line arguments.
- **Docker Compatibility**: Easily deployable within a Docker container.

## Prerequisites
- Go (1.15 or later)
- Docker (for containerized deployment)
- XIVLauncher configured to use the XL Authenticator app/OTP macro support

## Running the Application

### Locally
To run the application locally, follow these steps:

1. **Set up environment variables** (optional):
    - `XIVOTP_SECRET`: Your secret key for OTP generation, spaces removed.
    - `XIVLAUNCHER_IPS`: Single or multiple IP addresses. If multiple addresses are used, they should be comma-separated without spaces.
   ```
   export XIVOTP_SECRET="your_secret"
   export XIVLAUNCHER_IPS="ip1,ip2,..."
   ```

2. **Run the application**:
    - Using environment variables:
      ```
      go run main.go
      ```
    - Using command-line arguments:
      ```
      go run main.go <secret> <ip1> <ip2> ...
      ```

### In a Docker Container
To run the application in a Docker container, follow these steps:

1. **Build the Docker image**:
   ```
   docker build -t xivlauncher-otp-service .
   ```

2. **Run the container**:
    - Using environment variables:
      ```
      docker run -d --name xivlauncher-otp-service \
      -e XIVOTP_SECRET="your_secret" \
      -e XIVLAUNCHER_IPS="ip1,ip2,..." \
      xivlauncher-otp-service
      ```
    - Alternatively, pass command-line arguments directly (ensure the Dockerfile supports this method).

## Configuration
The application can be configured either through environment variables or command-line arguments:

- `XIVOTP_SECRET`: The secret key used for OTP generation. This is only available when you initially set up your Software Authenticator in your Square Enix Account center. If you already have this setup, simply remove it and re-add it. You can add it  to Authy or other applications as a backup at this time as well.
- `XIVLAUNCHER_IPS`: A comma-separated list of IP addresses where the XIVLauncher service is (or will be) running. This has been tested to work on multiple types of devices including Desktop PCs, Laptops, and handhelds such as the Steam Deck. Ensure that the IP addresses for these devices are static otherwise this service will need to be relaunched with the new IP addresses.

## Logging
The application logs the following events:
- Start-up information, including the IPs it will be monitoring.
- Successful OTP generation and submission.
- Critical errors.

Regular connection attempts and failures are not logged to avoid cluttering the output.
