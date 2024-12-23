"use client";

import Image from "next/image";
import Link from "next/link";
import { useEffect, useState } from "react";
import { FaChevronLeft, FaShoppingCart, FaStar } from "react-icons/fa";

const Panel = () => {
    const [productCategories, setProductCategories] = useState([])
    const [stores, setStores] = useState([])
    const [products, setProducts] = useState([])

    useEffect(() => {
        fetch("http://localhost/api/product-categories", {
            headers: {
                'Content-Type': 'application/json',
            }
        })
            .then(res => res.json())
            .then(({ data }) => {
                setProductCategories(data)
            })

        fetch("http://localhost/api/panel/stores/nearest", {
            headers: {
                'Content-Type': 'application/json',
                authorization: 'Bearer ' + localStorage.getItem("access_token")
            }
        })
            .then(res => res.json())
            .then(({ data }) => setStores(data))
            .catch((err) => {
                console.log("an error accured => ", err)
            })

        fetch("http://localhost/api/panel/products?sort=rate", {
            headers: {
                'Content-Type': 'application/json',
                authorization: 'Bearer ' + localStorage.getItem("access_token")
            }
        })
            .then(res => res.json())
            .then(({ data }) => setProducts(data))
            .catch((err) => {
                console.log("an error accured => ", err)
            })
    }, [])

    return <>
        <section className="mb-10">
            <h1 className="font-bold mb-5 text-xl">دسته بندی ها</h1>
            <section className="">
                {productCategories.length > 0 && (
                    <ul className="grid grid-cols-6 gap-6">
                        {productCategories.map(c => (
                            // TODO: change it to link
                            <li key={c.id} className="relative h-32 bg-white cursor-pointer rounded-lg overflow-hidden border-2 border-white shadow">
                                <Image src={c.image} fill={true} alt={c.name} />
                                <div className="bg-inherit pr-3 pl-4 py-1 absolute bottom-0 right-0 rounded-tl-xl inline-flex items-center">
                                    <span className="ml-2">{c.name}</span>
                                    <FaChevronLeft />
                                </div>
                            </li>
                        ))}
                    </ul>
                )}
            </section>
        </section>

        {stores.length > 0 && (
            <section className="mb-10">
                <h1 className="font-bold mb-5 text-xl">نزدیک ترین رستوران ها</h1>
                <div className="flex">
                    {stores.map(store => (
                        <Link href={`/stores/${store.slug}`} key={store.id} className="w-52 h-72 border ml-5 rounded-xl shadow overflow-hidden">
                            <section className="w-full h-1/3 relative mb-16">
                                <Image src={store?.store_type.image} fill={true} alt={store?.store_type.name} />
                                <section className="rounded-xl overflow-hidden absolute z-10 left-14" style={{ bottom: '-50px' }}>
                                    <Image
                                        src={store.logo}
                                        width={100}
                                        height={100}
                                        alt={store.name}
                                    />
                                </section>
                            </section>
                            <section className="flex flex-col items-center justify-center px-3">
                                <span className="font-bold mb-2">{store.name}</span>
                                <span className="inline-flex items-center text-xs mb-2">
                                    <FaStar className="text-yellow-400 ml-2" /> {"۴.۳ (۳.۵۰۰)"}
                                </span>
                                <span className="text-gray-700 text-sm">{store.store_type.name}</span>
                            </section>
                        </Link>
                    ))}
                </div>
            </section>
        )}

        {products.length > 0 && (
            <section className="mb-10">
                <h1 className="font-bold mb-5 text-xl">برترین غذاها</h1>
                <div className="flex">
                    {products.map(product => (
                        <Link href={`/products/${product.slug}`} key={product.id} className="w-60 h-80 border ml-5 rounded-xl shadow overflow-hidden py-5 px-2 relative">
                            <section className="w-full h-1/2 relative mb-3 rounded-lg overflow-hidden">
                                <Image
                                    src={product.image}
                                    fill={true}
                                    alt={product.name}
                                />
                            </section>
                            <section className="flex flex-col items-center justify-center px-3 h-auto">
                                <span className="font-bold mb-1 text-center">{product.name}</span>

                                <Link className="mb-3" href={`stores/${product?.store.slug}`}>
                                    <span className="text-xs">
                                        {product?.store.name}
                                    </span>
                                </Link>

                                <div className="text-xs flex justify-between w-full mb-4">
                                    <span>
                                        {product.price} تومان
                                    </span>
                                    <span className="inline-flex items-center">
                                        <FaStar className="text-yellow-400 ml-2" /> {"۴.۳"}
                                    </span>
                                </div>
                                <button className="w-11/12 bg-pink-700 text-sm rounded text-white py-2 px-3 absolute bottom-3 flex justify-between items-center">
                                    <span>
                                        افزودن به سبد خرید
                                    </span>
                                    <FaShoppingCart />
                                </button>
                            </section>
                        </Link>
                    ))}
                </div>
            </section>
        )}
    </>
};

export default Panel;
