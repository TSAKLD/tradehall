db:
	docker run -v steamsale:/var/lib/postgresql/data/ -p "5432:5432" -e POSTGRES_PASSWORD=asdbnm321 -e POSTGRES_USER=kr -e POSTGRES_DB=steamsale -d postgres:14.2