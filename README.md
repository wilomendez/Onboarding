# Onboarding

This repo consists in a serie of steps required for probe and acquire the basic concepts of some useful AWS services used in Ualá.

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

## Step 2 - Api Gateway :white_check_mark:

* Create another lambda to query a **Contact** via API, that complies with the flow: ``` API Gateway -> Lambda -> Dynamo DB ```.
* The next services must be able called using any tool (for example [Postman](https://www.postman.com/downloads/)).
* The services response must be JSON objects (See more ``` API Gateway -> Integration Response -> Mapping Templates ```)

**Requirements:**
* Contact creation API must be exposed using API Gateway (use lambda from step 1)
  * Route: ```<server>/contacts```
  * Method: **POST**

* Contact query API must be exposed using API Gatway
  * Route: ```<server>/contacts/{id}```
  * Method: **GET**

**Recommended readings:**
* API Gateway and Lambda integration:
  * [Create a simple microservice using Lambda and API Gateway](https://docs.aws.amazon.com/lambda/latest/dg/services-apigateway-blueprint.html).
  * [Integrating API with AWS services lambda](https://docs.aws.amazon.com/apigateway/latest/developerguide/integrating-api-with-aws-services-lambda.html#api-as-lambda-proxy-create-api-resources).
* API Gateway response mapping:
  * [Working with models and mapping templates](https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html).
* Mapping Example:
  * [Tutorial: Crear una API de REST importando un ejemplo](https://docs.aws.amazon.com/es_es/apigateway/latest/developerguide/api-gateway-create-api-from-example.html)



## Services

* **λ** : 
  * **create-contact-dev-wm-go**
  * **get-contact-dev-wm-go**
  * **notify-contact-dev-wm-go**
* **API Gateway**
  * **Onboarding Golang WM**
* **DynamoDB TableName** 
  * **Contacts_WM**
* **Notifications**
  * **Topics**
    * **ContactsTopic_WM**
