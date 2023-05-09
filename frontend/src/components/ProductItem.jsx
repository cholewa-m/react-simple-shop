import React from 'react';

const ProductItem = ({ product, onAddToCart }) => {
    const handleClick = () => {
        onAddToCart(product);
      };
      

  return (
    <div>
      <h3>{product.Name}</h3>
      <p>{product.Price} zł</p>
      <button onClick={handleClick}>Dodaj do koszyka</button>
    </div>
  );
};

export default ProductItem;
