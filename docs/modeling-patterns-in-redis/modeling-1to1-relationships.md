# Modeling 1 to 1 relationships

## Scenario

    - List view of electronics and then a detailed view of each selected Electronic
    - Best buy select item, then view details of item in popup view
        - detailed view shows all details such as multiple photos, description, manufacturer, dimensions, weight, and so on.

    - 1 to 1 relationships using SQL
        - In a relational database, you may create a table called products where each row holds just enough data to display the information in the list view
        - Then, you may create another table called product_details where each row holds the rest of the details
        - You would also need a product_images table, where you store all of the images for a product
        - You can see the entity relationship diagram in Picture 3.
