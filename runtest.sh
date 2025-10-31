echo "🚀 Construyendo la app y levantando contenedores..."

docker compose down -v 
# down -v → limpia contenedores y volúmenes.

docker compose up -d --build
# mejor para levantar todo en desarrollo

# docker compose build --no-cache
# build --no-cache → compila la app desde cero.
# docker compose up -d 
# up -d → levanta DB y API en segundo plano.

# Esperar que la API esté lista
echo "Esperando a que la API esté disponible..."
sleep 5  # ajustar según tiempo de arranque

echo "Ejecutando tests CRUD..."
echo " Creando 3 películas..."

# 1) Crea 3 películas 
curl -s -X POST http://localhost:8080/peliculas \
  -H "Content-Type: application/json" \
  -d '{"title":"Matrix","director":"Wachowski","year":1999,"genre":"Sci-Fi"}'
echo ""

curl -s -X POST http://localhost:8080/peliculas \
  -H "Content-Type: application/json" \
  -d '{"title":"Inception","director":"Christopher Nolan","year":2010,"genre":"Sci-Fi"}'
echo ""

curl -s -X POST http://localhost:8080/peliculas \
  -H "Content-Type: application/json" \
  -d '{"title":"The Godfather","director":"Francis Ford Coppola","year":1972,"genre":"Crime"}'
echo ""

echo "Actualizando película (ID=2)"
# 2) Actualizar una película (PUT). Se modifica los campos titulo y genero
curl -s -X PUT http://localhost:8080/peliculas/2 \
  -H "Content-Type: application/json" \
  -d '{"title":"Inception Updated","director":"Christopher Nolan","year":2010,"genre":"Thriller"}'
echo ""

echo "Obteniendo película por ID (ID=2)..."
# 3) Obtener una película por ID
curl -s http://localhost:8080/peliculas/2
echo ""

echo "Listando todas las películas..."
# 4) Listar todas las películas
curl -s http://localhost:8080/peliculas
echo ""

echo "Eliminando película (ID=1)..."
# 5) Eliminar una película
curl -s -X DELETE http://localhost:8080/peliculas/1
echo ""

echo "Listando todas las películas después de eliminar la primera..."
# 6) Volver a listar las películas
curl -s http://localhost:8080/peliculas
echo ""

echo "✅ CRUD completo ejecutado"

#Crear 3 películas
#Actualizar la película con ID 2
#Obtener la película actualizada por ID
#Listar todas las películas para ver cambios
#Eliminar la película con ID 1
#Listar todas las películas nuevamente para confirmar eliminación