This function is used to extract the quantity, unit, and product name from the given text. If the text does not contain quantity and unit, it will only return the product name.

You can obtain the results by sending a POST request as follows:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"text": "Tadım Fındık İçi 90 G"}' http://localhost:3000
```

It will return a JSON response like this:

```json
{
  "product_name": "tadım fındık içi",
  "quantity": 90,
  "unit": "g"
}
```

Note: The endpoint `http://localhost:3000` is used as an example in the cURL command. In a real scenario, you should replace it with the actual endpoint where the function is hosted.