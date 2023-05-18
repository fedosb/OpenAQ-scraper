# OpenAQ scraper
This is a sample project for scraping and analyzing data from OpenAQ API.

### Stack:
- `Go` for scraper
- `Jupyter Notebook` for data analysis and presenting
- `PostgreSQL` as DB (with `Adminer`)
- `Docker` and `docker-compose` for deployment
- `GitHub Actions` workflows for running code lint and unit tests

### How to run:
1. Run `docker-compose -f docker-compose.yml up --build -d` to deploy DB, adminer and scraper
2. Send `GET /api/v1/scraping` to scraper
3. Run all statements at `presenter.ipynb` to filter and analyze data

### Ports exposed:
- `5432` – PostgreSQL
- `5051` – Adminer
- `8081` – Scraper

### Scraper API:
- `/api/v1/scraping` – begin scraping data
- `/api/v1/measurements` – get measurements
- `/api/v1/measurements/cities` – get measurements cities
- `/api/v1/measurements/locations` – get measurements locations
- `/api/v1/measurements/parameters` – get measurements parameters
##### Query params:
- `country` (ex. US) – defaults to US
- `city` (ex. Columbus)
- `location` (ex. South Valley)
- `limit` (ex. 10) – defaults to 1000
- `count` (ex. 1000) – if passed 0, the maximum available will be chosen
- `parameter` (ex. no)
