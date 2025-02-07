package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Article struct {
	Title       string
	Link        string
	Description string
	SourceName  string `json:"source_name"`
}

type ApiResponse struct {
	Status  string
	Error   string
	Results []Article
}

func newError(err error) ApiResponse {
	return ApiResponse{
		Status: "error",
		Error:  err.Error(),
	}
}

func GetResults(category string) ApiResponse {
	// 	example := []byte(`
	// {
	//   "status": "success",
	//   "totalResults": 23111,
	//   "results": [
	//     {
	//       "article_id": "71a889992044d69300357342472042e7",
	//       "title": "El tiempo en Mogán: previsión meteorológica para hoy, martes 4 de febrero",
	//       "link": "https://www.laprovincia.es/tiempo/2025/02/04/tiempo-mogan-prevision-meteorologica-hoy-113978857.html",
	//       "keywords": null,
	//       "creator": [
	//         "La Provincia - Diario de Las Palmas"
	//       ],
	//       "video_url": null,
	//       "description": "El día de hoy, 4 de febrero de 2025, en Mogán, se espera un tiempo variable con intervalos nubosos y la posibilidad de lluvia escasa a lo largo de la jornada. Desde las primeras horas de la mañana, el cielo estará cubierto con nubes que podrían dejar caer algunas gotas, especialmente en los periodos de la madrugada hasta la mañana, donde la probabilidad de precipitación se sitúa en un 50%. A medida que avance el día, la situación mejorará ligeramente, con cielos más despejados y temperaturas que oscilarán entre los 15 y 20 grados .",
	//       "content": "ONLY AVAILABLE IN PAID PLANS",
	//       "pubDate": "2025-02-04 03:35:59",
	//       "pubDateTZ": "UTC",
	//       "image_url": "https://estaticos-cdn.prensaiberica.es/clip/bd16dd9f-d00b-408e-8afa-a2f59426492e_16-9-aspect-ratio_default_0.jpg",
	//       "source_id": "laprovincia",
	//       "source_priority": 563463,
	//       "source_name": "La Provincia",
	//       "source_url": "https://www.laprovincia.es",
	//       "source_icon": "https://i.bytvi.com/domain_icons/laprovincia.png",
	//       "language": "spanish",
	//       "country": [
	//         "spain"
	//       ],
	//       "category": [
	//         "top"
	//       ],
	//       "ai_tag": "ONLY AVAILABLE IN PROFESSIONAL AND CORPORATE PLANS",
	//       "sentiment": "ONLY AVAILABLE IN PROFESSIONAL AND CORPORATE PLANS",
	//       "sentiment_stats": "ONLY AVAILABLE IN PROFESSIONAL AND CORPORATE PLANS",
	//       "ai_region": "ONLY AVAILABLE IN CORPORATE PLANS",
	//       "ai_org": "ONLY AVAILABLE IN CORPORATE PLANS",
	//       "duplicate": false
	//     },
	//     {
	//       "article_id": "71a889992044d69300357342472042e7",
	//       "title": "El tiempo en Mogán: previsión meteorológica para hoy, martes 4 de febrero",
	//       "link": "https://www.laprovincia.es/tiempo/2025/02/04/tiempo-mogan-prevision-meteorologica-hoy-113978857.html",
	//       "keywords": null,
	//       "creator": [
	//         "La Provincia - Diario de Las Palmas"
	//       ],
	//       "video_url": null,
	//       "description": "El día de hoy, 4 de febrero de 2025, en Mogán, se espera un tiempo variable con intervalos nubosos y la posibilidad de lluvia escasa a lo largo de la jornada. Desde las primeras horas de la mañana, el cielo estará cubierto con nubes que podrían dejar caer algunas gotas, especialmente en los periodos de la madrugada hasta la mañana, donde la probabilidad de precipitación se sitúa en un 50%. A medida que avance el día, la situación mejorará ligeramente, con cielos más despejados y temperaturas que oscilarán entre los 15 y 20 grados .",
	//       "content": "ONLY AVAILABLE IN PAID PLANS",
	//       "pubDate": "2025-02-04 03:35:59",
	//       "pubDateTZ": "UTC",
	//       "image_url": "https://estaticos-cdn.prensaiberica.es/clip/bd16dd9f-d00b-408e-8afa-a2f59426492e_16-9-aspect-ratio_default_0.jpg",
	//       "source_id": "laprovincia",
	//       "source_priority": 563463,
	//       "source_name": "La Provincia",
	//       "source_url": "https://www.laprovincia.es",
	//       "source_icon": "https://i.bytvi.com/domain_icons/laprovincia.png",
	//       "language": "spanish",
	//       "country": [
	//         "spain"
	//       ],
	//       "category": [
	//         "top"
	//       ],
	//       "ai_tag": "ONLY AVAILABLE IN PROFESSIONAL AND CORPORATE PLANS",
	//       "sentiment": "ONLY AVAILABLE IN PROFESSIONAL AND CORPORATE PLANS",
	//       "sentiment_stats": "ONLY AVAILABLE IN PROFESSIONAL AND CORPORATE PLANS",
	//       "ai_region": "ONLY AVAILABLE IN CORPORATE PLANS",
	//       "ai_org": "ONLY AVAILABLE IN CORPORATE PLANS",
	//       "duplicate": false
	//     }
	//   ],
	//   "nextPage": "1738640069077386592"
	// }`)

	url := fmt.Sprintf("https://newsdata.io/api/1/news?apikey=%s&country=es&language=es&category=%s", os.Getenv("API_KEY"), strings.ToLower(category))
	res, err := http.Get(url)
	if err != nil {
		return newError(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		newError(fmt.Errorf("Status code %v", res.StatusCode))
	}
	if err != nil {
		newError(err)
	}

	var apiRes ApiResponse

	if err := json.Unmarshal(body, &apiRes); err != nil {
		return newError(err)
	}

	return apiRes
}
