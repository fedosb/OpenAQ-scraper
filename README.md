# TPBDM
Technology platform for big data management

### Stack:
- `Go` for scraper
- `PostgreSQL` as DB
- `Jupyter Notebook` for data analysis and presenting
- `Docker` and `docker-compose` for deployment
- `GitHub Actions` workflows for running code lint and unit tests

### How to run:
1. Run `docker-compose -f docker-compose.yml up --build -d` to deploy DB, adminer and scraper
2. Send `GET localhost:8081/api/v1/scraping` with following params:
   - `country` (ex. US) – defaults to US
   - `city` (ex. Columbus)
   - `location` (ex. South Valley)
   - `limit` (ex. 10) – defaults to 1000
   - `count` (ex. 1000) – if passed 0, the maximum available will be chosen
   - `parameter` (ex. no)
3. Run all statements at `presenter.ipynb` to filter and analyze data
