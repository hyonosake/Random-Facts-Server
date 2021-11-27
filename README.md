# Random-Facts-Server

#### Random Facts is a RESTful API service consisting of PostgreSQL database that stores random facts and golang API service that handles POST, GET and PUT requests.

# API Description
## 1. Recieve random fact
#### Query: Header ```GET URL/fact ```
###### Response on success:
```
{
    "title": "Lorem Ipsum",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry.
    Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown
    printer took a galley of type and scrambled it to make a type specimen book."
    "links": [
        "Loremipsum.com/Lorem"
    ]
}
```
###### Response on error:

## 2. Recieve specific fact
#### Query: Header ```GET URL/fact/$passed_id ```
###### Response on success:
```
{
    "id": $passed_id
    "title": "Lorem Ipsum",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry.
    Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown
    printer took a galley of type and scrambled it to make a type specimen book."
    "links": [
        "Loremipsum.com/Lorem"
    ]
}
```
###### Response on error:


## 2. Post single/multiple facts
#### Query:
###### Header ```POST URL/fact ```
###### Body
```
{
  "facts": [
      {
          "title": "Lorem Ipsum",
          "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry.
          Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown
          printer took a galley of type and scrambled it to make a type specimen book."
          "links": [
          "Loremipsum.com/Lorem"
          ]
      },
      {
          "title": "Lorem Ipsum",
          "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry.
          Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown
          printer took a galley of type and scrambled it to make a type specimen book."
          "links": [
          "Loremipsum.com/Lorem"
          ]
    },
    {
        "title": "Lorem",
        "description": "Ipsum"
      }
  ]
}
```
###### Response on success:
```

```
###### Response on error:




