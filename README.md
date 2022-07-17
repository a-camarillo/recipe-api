# recipe-api
A simple REST API for recipes

This API contains two endpoints:

`/recipes/`

`/ingredients/`

### Use the API

If you would like to mess around with the API just make sure you have Docker installed on your system.

1. Clone this repository to your local machine and `cd recipe-api`
2. Make sure you don't have anything running on `PORT 8000` 
  * If you have an important service running on `PORT 8000` you will have to edit the `docker-compose.yml` file
  * Change 
    ```  
    ports: 
    - "8000:8000"
    ```
    to 
    ```
    ports:
    - "8000:{YOUR_PORT_OF_CHOICE}"
    ```
    this will expose `PORT 8000` within the container to whichever port you choose on your local machine
3. Build the service with the command `docker compose up` 
4. In a browser navigate to `localhost:8000/api/`or whichever port you decided to edit to. If you see the default Django Rest Framework UI and two endpoints then everything should be working and you should be able to interact with the API. Alternatively, once the Django server is up and running you can use a program like Postman to interact with the API.
