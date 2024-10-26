# 2024/03/29

Initial version.

# [v1.0.0] - 2024/10/26

## Fixes
- **High Memory Usage**: Investigated and resolved excessive memory consumption observed after multiple WebSocket sessions.
- **Cross-Site Request Forgery**: Fixed reported cross-site request forgery vulnerabilities to enhance security.
- **WebSocket Message Count**: Corrected the server's miscounting of messages in WebSocket session statistics.

## Enhancements
- **API Enhancement**: Extended the API to return hex values in WebSocket responses.
- **Command Line Client**: Created a command line client as a separate executable for easier interaction with the application.
- **Hex Value Generation**: Enhanced the random string generator to produce only hex values.

### Other Changes
- **Go Upgrade**: Upgraded Go version to **1.23.2**, addressing known vulnerabilities and improving security.