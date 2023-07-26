
package main

import (
    "encoding/json"
    "log"
    "github.com/gofiber/fiber/v2"
    "regexp"
    "strconv"
    "strings"
    "fmt"
)


func extractQuantityAndUnit(text string) map[string]interface{} {
	// Regex pattern to match numbers with decimal points or commas (e.g., 1.5, 1,000, 0,5)
	re := regexp.MustCompile(`\d+([.,]\d+)?`)

	units := map[string]bool{
		"g":  true,
		"gr": true,
		"kg": true,
		"ml": true,
		"cl": true,
		"l":  true,
		"lt": true,
	}

	tokens := strings.Split(text, " ")
	quantity := "nan"
	value := "nan"

	for i, token := range tokens {
		if re.MatchString(token) && i+1 < len(tokens) && units[strings.ToLower(tokens[i+1])] {
			quantity = strings.ToLower(tokens[i+1])
			value = strings.Replace(token, ",", ".", -1)
			break
		} else if re.MatchString(token) && i+1 < len(tokens) && units[strings.ToLower(tokens[i+1])] {
			quantity = strings.ToLower(tokens[i+1])
			value = strings.Replace(token, ",", ".", -1)
			break
		} else if units[strings.ToLower(token)] {
			quantity = strings.ToLower(token)
			break
		}
	}

	// Convert value to float64
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		parsedValue = 0.0
	}

	if quantity == "nan" && value == "nan" {
		return map[string]interface{}{
			"miktar":      parsedValue,
			"birim":       quantity,
			"product_name": text,
		}
	}

	// Convert text to lowercase and remove quantity and value from it
	text = strings.ToLower(text)
	if quantity != "nan" {
		text = strings.Replace(text, quantity, "", -1)
	}
	if value != "nan" {
		text = strings.Replace(text, strconv.Itoa(int(parsedValue)), "", -1)
		text = strings.Replace(text, strconv.FormatFloat(parsedValue, 'f', -1, 64), "", -1)
	}

	// Trim leading and trailing spaces
	text = strings.TrimSpace(text)

	return map[string]interface{}{
		"quantity":      parsedValue,
		"unit":       quantity,
		"product_name": text,
	}
}


func main() {
    app := fiber.New()

    app.Post("/", func(c *fiber.Ctx) error {
        // JSON verisinden gelen metni alın
        var requestData struct {
            Text string `json:"text"`
        }
        if err := c.BodyParser(&requestData); err != nil {
            return c.SendStatus(fiber.StatusBadRequest)
        }

        result := extractQuantityAndUnit(requestData.Text)
        jsonString, err := json.Marshal(result)
		fmt.Printf("Hllo")
        if err != nil {
            return c.SendString("Sonuç serileştirilirken bir hata oluştu")
        }
        return c.SendString(string(jsonString))
    })

    log.Fatal(app.Listen(":3000"))
}

