import { useEffect, useState } from 'react';


const Cart = ({ items }) => {
    const [newItems, setNewItems] = useState(items)
    const [allPrice, setAllPrice] = useState(0);

const clear = () => {
    window.localStorage.clear()
    setNewItems([])
    setAllPrice(0)
   }

   useEffect(()=>{
    const sum = items.reduce((acc, item) => acc + parseFloat(item.Price), 0)
    setAllPrice(sum)
   },[items])
  return (
    <div>
      <h1>Koszyk</h1>
      <ul>
      {newItems.map((item) => (
  <li key={item.ID}>{item.Name} - {item.Price} zł</li>
))}

      </ul>
      <div>Suma: {allPrice}</div>
      <p><a href='/payments'>Przejdź do płatności</a></p>
      <button onClick={clear}>Wyczyść koszyk</button>
    </div>
  );
};

export default Cart;
