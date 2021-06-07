## REST API (Fiber) + MySQL

### Функционал:

```
* POST    http://localhost:8000/api/contacts     добавление контакта через JSON {"name": "yourName", "number": "yourNumber"}
* GET     http://localhost:8000/api/contacts     получение списка контактов
* DELETE  http://localhost:8000/api/contacts/id  удаление контакта по id
* GET     http://localhost:8000/api/contacts/id  получение информации контакта по id
* PUT     http://localhost:8000/api/contacts/id  обновление контакта по id через JSON {"name": "yourName", "number": "yourNumber"}
```
