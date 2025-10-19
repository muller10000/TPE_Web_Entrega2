echo "probando metodo POST"

curl -X POST http://localhost:8080/peliculas \
  -H "Content-Type: application/json" \
  -d '{"title":"Matrix","director":"Wachowski","year":1999,"genre":"Sci-Fi"}'

echo "probando metodo GET"  

# Listar películas
curl http://localhost:8080/peliculas
