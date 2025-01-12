import Image from "next/image";
import { useState } from "react";
import { FaAngleDown } from "react-icons/fa";
import { convertEnToFaNumber } from "../utils/number";

const Order = ({ order, rerender }) => {
    const [showItems, setShowItems] = useState(false);

    const handlePay = () => {
        fetch(`http://localhost/api/profile/orders/${order.id}/pay`, {
            headers: {
                authorization: "Bearer " + localStorage.getItem("access_token"),
            },
        })
            .then((res) => res.json())
            .then(() => rerender());
    };

    const handleToggleShowItems = () => {
        setShowItems((p) => !p);
    };

    return (
        <li className="mb-4">
            <section className="flex">
                <Image
                    src={order.store.logo}
                    width={70}
                    height={70}
                    alt={order.store.name}
                />
                <section>
                    <h3 className="font-bold">{order.store.name}</h3>
                    <small>
                        {new Date(order.created_at).toLocaleDateString("fa-IR", {
                            day: "numeric",
                            weekday: "long",
                            month: "long",
                            hour: "2-digit",
                            minute: "2-digit",
                        })}
                    </small>
                </section>
            </section>

            <section className="py-4 grid grid-cols-2">
                {order.status_label === "pending" ? (
                    <>
                        <button
                            className="bg-pink-600 text-white rounded py-2 mx-1"
                            onClick={handlePay}
                        >
                            پرداخت
                        </button>
                        <button className="border-2 rounded py-2">انصراف</button>
                    </>
                ) : (
                    <span>{order.status_label_fa}</span>
                )}
            </section>

            {order.status_label === "done" && (
                <section className="border">
                    <button
                        className="py-1 w-full inline-flex justify-center items-center"
                        onClick={handleToggleShowItems}
                    >
                        مشاهده فاکتور <FaAngleDown />
                    </button>

                    {showItems && (
                        <ul className="px-3 py-2">
                            {order.items.map((item) => (
                                <li
                                    className="flex justify-between text-sm border-b py-2"
                                    key={item.id}
                                >
                                    <span>{item.product.name}</span>
                                    <span>
                                        {item.quantity} * {convertEnToFaNumber(item.price)} تومان
                                    </span>
                                </li>
                            ))}

                            <li className="flex justify-between text-sm border-b py-2">
                                <span>مجموع</span>
                                <span>{convertEnToFaNumber(order.amount)} تومان</span>
                            </li>
                        </ul>
                    )}
                </section>
            )}
        </li>
    );
};

export default Order;
