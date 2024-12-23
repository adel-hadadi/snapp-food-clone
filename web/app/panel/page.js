"use client";

import Image from "next/image";
import Link from "next/link";
import { useEffect, useState } from "react";
import { FaChevronLeft } from "react-icons/fa";

const Panel = () => {
    const [productCategories, setProductCategories] = useState([])
    const [stores, setStores] = useState([])

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

        fetch("http://localhost/api/stores", {
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then(res => res.json())
            .then(({ data }) => setStores(data))
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

        <section>
            <h1 className="font-bold mb-5 text-xl">نزدیک ترین رستوران ها</h1>
            <div className="flex">
                {stores.length > 0 && stores.map(store => (
                    <Link href={`/stores/${store.slug}`} key={store.id} className="w-52 h-72 border ml-5">
                        <section className="w-full h-1/2 relative">
                            <Image src={store?.store_type.image} fill={true} alt={store?.store_type.name} />
                        </section>
                        <section className="text-center">
                            <span className="font-bold">{store.name}</span>
                        </section>
                    </Link>
                ))}
            </div>

        </section>
    </>
};

export default Panel;
