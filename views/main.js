const addToCartButtons = document.querySelectorAll('.add-to-cart');

addToCartButtons.forEach(button => {
	button.addEventListener('click', () => {
		const productId = button.getAttribute('data-id');
		addToCart(productId);
	});
});

function addToCart(productId) {
	// Implementasi untuk menambahkan produk ke keranjang belanja
	console.log(`Produk dengan ID ${productId} ditambahkan ke keranjang belanja.`);
}