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

**λ** : **create-contact-dev-wm-go**
**DynamoDB TableName** : **Contacts_WM**

## Step 1B - Lambda :white_check_mark:

* The above lambda will need to be configured for the Alias.

**TEST**. Configure this Alias with any other version except `$LATEST`.

**Recommended resources for Activity 1**

* **JAVA**
  * [Lambdas and Dynamo tables tutorial](http://www.baeldung.com/aws-lambda-dynamodb-java)
* **GO**
  * [Golang Lambdas](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html).
  * [Packaging Lambdas in Golang](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html).
  * [DynamoDB](https://antklim.medium.com/dynamodb-expressions-and-go-b8230c253e1f)
* [Lambda Alias and Versions](https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html)

