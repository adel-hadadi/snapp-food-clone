"use client";

import { FaShoppingCart } from "react-icons/fa";
import Modal from "../components/modal";
import { useCallback, useEffect, useState } from "react";
import Order from "../components/order";

const Cart = () => {
  const [showCartModal, setShowCartModal] = useState(false);
  const [orders, setOrders] = useState([]);
    const [, updateState] = useState();
    const forceUpdate = useCallback(() => updateState(), [])

  useEffect(() => {
    fetch("http://localhost/api/profile/orders", {
      headers: {
        authorization: "Bearer " + localStorage.getItem("access_token"),
      },
    })
      .then((res) => res.json())
      .then(({ data }) => {
        setOrders(data);
      });
  }, []);

  return (
    <>
      <button
        className="flex items-center px-5 py-3 text-sm"
        onClick={() => setShowCartModal(true)}
      >
        <FaShoppingCart />
        <small className="mr-1">سفارشات</small>
      </button>

      <Modal
        isOpen={showCartModal}
        handleToggleModal={() => setShowCartModal((p) => !p)}
        title="لیست سفارشات"
      >
        <ul>
          {orders.length > 0 &&
            orders.map((order) => <Order order={order} key={order.id} rerender={forceUpdate} />)}
        </ul>
      </Modal>
    </>
  );
};

export default Cart;
