import React, { useEffect, useState } from 'react';
import ProductItem from './ProductItem';

const Products = ({ onAddToCart }) => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    async function fetchProducts() {
      try {
        const response = await fetch('http://localhost:1323/products');
        const data = await response.json();
        setProducts(data);
      } catch (error) {
        console.error('Error:', error);
      }
    }

    fetchProducts();
  }, []);

  return (
    <div>
      <h1>Produkty</h1>
      {products.map((product) => (
        <ProductItem key={product.ID} product={product} onAddToCart={() => onAddToCart(product)} />
      ))}
    <p><a href='/cart'>Przejdź do koszyka</a></p>
    </div>
  );
};

export default Products;
