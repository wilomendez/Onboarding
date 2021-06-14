# Onboarding

This repo consists in a serie of steps required for probe and acquire the basic concepts of some useful AWS services used in USalá

## Requirements
### Step 1A :white_check_mark:

Create a Lambda to insert a contact that executes this flow `Lambda -> DynamoDB`.

* The service must be processed for a lambda function and executed in this case by **Golang Programming language**.
* Contact must be stored in a DynamoDB table called `Contacts` (or similar names that not makes conflicts with any other existent projects) and follow this aspect:

````
Table: Contacts
         id: String (key)
         firstName: String
         lastName: String
         status: String
````

Contact must be stored with **CREATED** status.


