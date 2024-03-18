# Weather Status

This project is a simple web application that displays weather status based on water and wind values obtained from a JSON file served by a Go server. The status includes whether the conditions are safe, cautionary, or dangerous based on predefined thresholds for water and wind values.

## Technologies Used
- Go (Golang): For server-side programming to serve the JSON file.
- HTML/JavaScript: For client-side programming to create the web interface and update data dynamically.

## Installation and Setup
1. Clone the repository:

    ```
    git clone https://github.com/hakimasyrofi/Hacktiv_DTS_Golang.git
    ```

2. Navigate to the project directory:

    ```
    cd assignment3
    ```

3. Run the Go server:

    ```
    go run main.go
    ```

4. Open your web browser and go to [http://localhost:8080](http://localhost:8080) to view the weather status dashboard.

## Usage
- The web page will display the current water and wind values obtained from the JSON file served by the Go server.
- The status will be automatically updated based on the following criteria:
    - Water:
        - Below 5 meters: Safe
        - Between 6 to 8 meters: Caution
        - Above 8 meters: Dangerous
    - Wind:
        - Below 6 meters per second: Safe
        - Between 7 to 15 meters per second: Caution
        - Above 15 meters per second: Dangerous
