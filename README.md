# Random-Facts-Server

#### Random Facts is a RESTful API service consisting of PostgreSQL database that stores random facts and golang API service that handles POST, GET and PUT requests.

# API Description
## 1. Recieve random fact
#### Query: Header ```GET URL/fact ```
#### Body: ignored
#### Response on success:
```
{
    "title": "Lorem Ipsum",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry."
    "links": [
        "Loremipsum.com/Lorem"
    ]
}
```
#### Response on error:
Consists of ```404 (Not Found)``` Header and JSON body:
```
{
    "status": 400,
    "message": "Some error message"
}
```

## 2. Recieve specific fact
#### Query: ```GET URL/fact/$passed_id ```
#### Body: ignored
#### Response on success:
Consists of ``` 302 (Found)``` Header and JSON body:
```
{
    "id": $passed_id
    "title": "Lorem Ipsum",
    "description": "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
        when an unknownprinter took a galley of type and scrambled it to make a type specimen book."
    "links": [
        "Lorem.com/Lorem",
        "Ipsum.net/Ipsum"
    ]
}
```
#### Response on error:
Consists of ```404 (Not Found)``` Header and JSON body:
```
{
    "status": 404,
    "message": "Some error message"
}
```
## 3. Post single/multiple facts
#### Query: ```POST URL/fact ```
#### Body
```
{
    "facts": [
        {
            "title": "A diabolo, qui est simia dei",
            "description": "English equivalent: Where god has a church the devil will have his chapel.",
            "links": [
                "https://en.wikiquote.org/wiki/Latin_proverbs"
            ]
        },
        {
            "title": "Some facts may not include Links",
            "description": "And you should be fine with that.",
        }
    ]
}
```
#### Response on success:
Consists of ```201 (Created)``` Header and JSON body of inserted IDs:
```
{
    "ids": [
        7,
        8,
        9
    ]
}
```
#### Response on error:
Consists of ```400 (Bad Request)``` Header and JSON body:
```
{
    "status": 400,
    "message": "Unable to POST error message"
}
```
## 4. Change info about specific fact
#### Query:  ```PUT URL/fact/42 ```
#### Body:
```
{
    "id": 42,
    "title": "Abyssus abyssum invocat.",
    "description": "English equivalent: Deep calls to deep. From the Bible, Psalm 42:7.",
    "links": [
        "https://en.wikiquote.org/wiki/Latin_proverbs",
        "https://www.bible.com/en-GB/bible/116/psa.42.7"
    ]
}
```
#### Response on success:
Consists of  ```302 (Found)``` Header and empty JSON body:
```
{}
```
#### Response on error:
Consists of ```400 (Bad ERequest)``` Header and JSON body:
```
{
    "status": 400,
    "message": "Some error message"
}
```
# To launch service locally:
1. git clone this repo
```
git clone https://github.com/hyonosake/Random-Facts-Server/ && cd Random-Facts-Server/
```
2. Use docker-compose to build containers with PostgreSQL and API service
```
docker-compose up
```
3. Open your browser or API-testing app, enter ```localhost:8080/fact``` and make some requests!
## Testing
tests folder provides json-exported Queries from Postman. 




