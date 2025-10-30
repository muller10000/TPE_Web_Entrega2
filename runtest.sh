echo "🚀 Construyendo la app y levantando contenedores..."

docker-compose down -v  # Baja contenedores y borra volúmenes temporales (DB limpia)
docker build --no-cache -t peliculas-api .
docker-compose up -d   # Levanta DB y API en segundo plano

# Esperar que la API esté lista
echo "Esperando a que la API esté disponible..."
sleep 5  # ajusta según tu tiempo de arranque

echo "Ejecutando tests CRUD..."

# POST de prueba
curl -s -X POST http://localhost:8080/peliculas \
  -H "Content-Type: application/json" \
  -d '{"title":"Matrix","director":"Wachowski","year":1999,"genre":"Sci-Fi"}'
echo ""

# GET de prueba
curl -s http://localhost:8080/peliculas
echo ""

echo "✅ Tests completados"