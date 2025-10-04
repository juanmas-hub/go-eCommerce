# go-eCommerce

User: representa a una persona registrada en la aplicación.

Product: un artículo disponible para la venta en el catálogo.

Category: agrupa productos bajo un mismo tipo o temática (ej: ropa, tecnología).

Cart: el carrito de compras activo de un usuario donde guarda productos antes de comprar.

CartItem: un ítem dentro del carrito que indica qué producto y cuántas unidades tiene.

Order: un pedido confirmado con varios items de un usuario.

OrderItem: cada producto dentro de un pedido, con su cantidad y precio.

Relaciones:
User (1) -> Cart (1) -> CartItems (N) -> Product (1)

User (1) -> Orders (N) -> OrderItems (N) -> Product (1)

Category (1) -> Products (N)
