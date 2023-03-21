# Challenge: Cupón de compra
Aplicación construida en go y clean arquitecture el cual permite hacer uso del coupon para realizar la comprar tantos ítems marcados como favoritos sea posible.

## Redis DataBase
	1.- Tener instalado redis como manejador de base de datos.
	2.- La cadena de conexion por defecto es: localhost:6379
    3.- Si desea camnbiar dicha cadena de conexión, debe hacerlo en el archivo main.go

## Ejecución (Run)
	1.- Para levantar o ejecutar la app, hacerlo bajo el siguente código: go run main.go

## Depuración (Debugger)
	1.- Si desea realizar depuración en tiempo de ejcución, debes instalar a nivel de vscode el complemento GO
        https://marketplace.visualstudio.com/items?itemName=golang.Go
    2.- Ya se encuentra configurado launch.json para hacer depuraciones.
    3.- El archivo launch.json se encuentra localizado en el folder .vscode

## URLs (EndPoints)
_1.- Coupon_

_[GET] host:port/coupon | Example: http://localhost:8000/coupon

	"Body Params" (Ejemplo)
	{
	  "item_ids":  ["MCO899299522", "MCO921554676", "MCO878860564", "MCO854851024", "MCO952084987", "MCO571817761", "MCO-1171704643", "MCO607792825","MCO1084100492"],
	  "amount": 200000.00
	}

	"Result" (Ejemplo)
	{
	    "item_ids": [
            {
                "id": "MCO952084987",
                "site_id": "MCO",
                "title": "Tarjeta De Memoria Adata Ausdx128guicl10a1-ra1  Premier Con Adaptador Sd 128gb",
                "price": 56900,
                "seller_id": 170421130
            },
            {
                "id": "MCO899299522",
                "site_id": "MCO",
                "title": "Disco Sólido Ssd Interno Kingston Sa400s37/240g 240gb Negro",
                "price": 141360,
                "seller_id": 170421130
            }
        ],
        "total": 198260
	}

_2.- Top_

_[GET] host:port/top | Example: http://localhost:8000/top

	Devuelve el Top5 de productos mas favoritos.
    Puede observar la cantidad de apariciones a traves de la propiedad (quantity) expuesta en el resultado.

	"Result" (Example)
	{
        {
            "id": "MCO952084987",
            "site_id": "MCO",
            "title": "Tarjeta De Memoria Adata Ausdx128guicl10a1-ra1  Premier Con Adaptador Sd 128gb",
            "price": 56900,
            "seller_id": 170421130,
            "quantity": 2
        },
        {
            "id": "MCO899299522",
            "site_id": "MCO",
            "title": "Disco Sólido Ssd Interno Kingston Sa400s37/240g 240gb Negro",
            "price": 141360,
            "seller_id": 170421130,
            "quantity": 1
        }
	}