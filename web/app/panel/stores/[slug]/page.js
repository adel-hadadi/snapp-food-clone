"use client";

import AxiosInstance from "@/app/utils/axiosInstance";
import { convertEnToFaNumber } from "@/app/utils/number";
import Image from "next/image";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { FaTrash } from "react-icons/fa";
import { IoStar } from "react-icons/io5";
import { toast, ToastContainer } from "react-toastify";

export default function Page() {
    const { slug } = useParams();
    const [store, setStore] = useState({ categories: [] });
    const [cart, setCart] = useState({ items: [], price: 0 });

    const handleAddProductToCart = (product) => {
        setCart((prev) => {
            return {
                price: prev.price + product.price,
                items: [...prev.items, { ...product, quantity: 1 }],
            };
        });
    };

    const handleIncrQuantity = (productId) => {
        const index = cart.items.findIndex((item) => item.id == productId);
        setCart((preventItems) => {
            return {
                ...preventItems,
                price: preventItems.price + preventItems.items[index].price,
                items: preventItems.items.map((item, i) =>
                    i === index ? { ...item, quantity: item.quantity + 1 } : item,
                ),
            };
        });
    };

    const handleDcrQuantity = (productId) => {
        const index = cart.items.findIndex((item) => item.id == productId);
        setCart((preventItems) => {
            return {
                ...preventItems,
                price: preventItems.price - preventItems.items[index].price,
                items: preventItems.items.map((item, i) =>
                    i === index ? { ...item, quantity: item.quantity - 1 } : item,
                ),
            };
        });
    };

    const handleDeleteItem = (productId) => {
        const product = cart.items.find((v) => v.id == productId);

        setCart((preventItems) => {
            return {
                ...preventItems,
                price: preventItems.price - product.price,
                items: preventItems.items.filter((item) => item.id !== productId),
            };
        });
    };

    const handleDeleteAllItems = () => {
        setCart({ price: 0, items: [] });
    };

    const handleSubmitOrder = async () => {
        const items = cart.items.map((item) => {
            return {
                product_id: item.id,
                quantity: item.quantity,
            };
        });

        const payload = {
            store_slug: slug,
            items: items,
        };
        const {data} = await AxiosInstance.post('/profile/orders', payload)
        if (data.success) {
            toast('سفارش با موفقیت ثبت شد')
        }
        setCart({});
    };

    useEffect(() => {
        fetch(`http://localhost/api/stores/${slug}`)
            .then((res) => res.json())
            .then(({ data }) => setStore(data));
    }, [slug]);

    return (
        <div className="grid gap-5 grid-cols-4">
            <section className="shadow rounded-xl py-6 px-5">
                <section className="flex">
                    <Image src={store.logo} alt={store.name} width={100} height={100} />
                    <section>
                        <span className="inline-flex justify-center">
                            <IoStar color="yellow" className="ml-1" />
                            <small>{store.rate}</small>
                        </span>
                        <h1>{store.name}</h1>
                    </section>
                </section>
            </section>

            <section className="col-span-2">
                {store?.categories.length > 0 &&
                    store?.categories.map((category) => (
                        <section className="border rounded-xl py-6 px-5" key={category.id}>
                            <h2 className="text-center border-b border-slate-300 pb-2 mb-4">
                                {category.name}
                            </h2>
                            <section className="grid grid-cols-2">
                                {category.products.map((product) => (
                                    <section key={product.id} className="py-3 px-2 border h-50">
                                        <section className="grid gap-2 grid-cols-3 h-32 mb-5">
                                            <section className="col-span-2 py-2">
                                                <h2 className="mb-2">{product.name}</h2>
                                                <p className="text-xs text-slate-500">
                                                    {product.description}
                                                </p>
                                            </section>
                                            <section className="relative">
                                                <Image
                                                    src={product.image}
                                                    alt={product.name}
                                                    fill={true}
                                                    className="rounded-xl"
                                                />
                                            </section>
                                        </section>
                                        <section className="flex items-center justify-between">
                                            <span className="text-sm">
                                                {convertEnToFaNumber(product.price)} تومان
                                            </span>
                                            <section className="flex items-center">
                                                <button
                                                    className="w-24 py-1 shadow text-pink-600 hover:text-white hover:bg-pink-600 rounded-full transition"
                                                    onClick={() => handleAddProductToCart(product)}
                                                >
                                                    افزودن
                                                </button>
                                                {/* 
                        <button className="w-9 h-9 ml-6 shadow text-pink-600 hover:text-white hover:bg-pink-600 rounded-xl transition">
                          +
                        </button>
                        <span>{item.quantity}</span>
                        <button className="w-9 h-9 mr-6 shadow text-pink-600 hover:text-white hover:bg-pink-600 rounded-xl transition flex items-center justify-center">
                          {item.quantity == 1 ? <FaTrash /> : "-"}
                        </button> */}
                                            </section>
                                        </section>
                                    </section>
                                ))}
                            </section>
                        </section>
                    ))}
            </section>

            <section>
                {cart.items.length > 0 && (
                    <>
                        <section className="border rounded-xl py-5 px-3 mb-3">
                            <section className="w-full flex justify-between items-center mb-6">
                                <h3 className="">سبد خرید ({cart.items.length})</h3>
                                <button
                                    className="text-pink-700 shadow w-7 h-7 rounded-full text-xs flex justify-center items-center"
                                    onClick={handleDeleteAllItems}
                                >
                                    <FaTrash />
                                </button>
                            </section>

                            <ul className="py-4 border-b">
                                {cart.items.map((item) => {
                                    return (
                                        <li key={item.id}>
                                            <p className="mb-2">{item.name}</p>
                                            <section className="flex items-center justify-between">
                                                <small>{convertEnToFaNumber(item.price)} تومان</small>

                                                <section className="flex items-center">
                                                    {item.quantity == 1 ? (
                                                        <button
                                                            className="w-9 h-9 ml-6 shadow text-slate-500 rounded-xl transition flex items-center justify-center"
                                                            onClick={() => handleDeleteItem(item.id)}
                                                        >
                                                            <FaTrash />
                                                        </button>
                                                    ) : (
                                                        <button
                                                            className="w-9 h-9 ml-6 shadow text-pink-600 hover:text-white hover:bg-pink-600 rounded-xl transition flex items-center justify-center"
                                                            onClick={() => handleDcrQuantity(item.id)}
                                                        >
                                                            -
                                                        </button>
                                                    )}
                                                    <span>{item.quantity}</span>
                                                    <button
                                                        className="w-9 h-9 mr-6 shadow text-pink-600 hover:text-white hover:bg-pink-600 rounded-xl transition"
                                                        onClick={() => handleIncrQuantity(item.id)}
                                                    >
                                                        +
                                                    </button>
                                                </section>
                                            </section>
                                        </li>
                                    );
                                })}
                            </ul>
                            <ul className="py-3 px-2 text-sm">
                                <li className="flex justify-between">
                                    <span>مجموع </span>{" "}
                                    <span>{convertEnToFaNumber(cart.price)} تومان</span>
                                </li>
                            </ul>
                        </section>
                        <button
                            className="bg-pink-600 hover:bg-pink-500 transition text-white rounded-lg w-full py-2"
                            onClick={handleSubmitOrder}
                        >
                            ثبت سفارش
                        </button>
                    </>
                )}
            </section>
        </div>
    );
}
